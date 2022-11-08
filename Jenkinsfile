properties([pipelineTriggers([githubPush()])])
pipeline {
  triggers {
    // Only trigger the cron build if on main branch and 10 AM Monday
    cron(env.BRANCH_NAME == 'main' ? 'TZ=GMT\n0 10 * * 1' : '')
  }

  options {
    disableConcurrentBuilds()
    buildDiscarder(logRotator(numToKeepStr: '3'))
    timeout(time: 30, unit: 'MINUTES')
  }

  agent {
    kubernetes {
      inheritFrom 'default'

      containerTemplates([
        containerTemplate(name: 'node', image: 'docker:20', resourceRequestCpu: '1', resourceRequestMemory: '2Gi', command: 'cat', ttyEnabled: true),
        containerTemplate(name: 'play', image: 'docker:20', resourceRequestCpu: '1', resourceRequestMemory: '2Gi', command: 'cat', ttyEnabled: true),
        containerTemplate(name: 'play-builder', image: 'flowdocker/play_builder:latest-java13', resourceRequestCpu: '1', resourceRequestMemory: '2Gi', command: 'cat', ttyEnabled: true)

      ])
    }
  }

  stages {
    stage('Checkout') {
      steps {
        checkoutWithTags scm

        script {
          VERSION = new flowSemver().calculateSemver() //requires checkout
        }
      }
    }

    stage('Commit SemVer tag') {
      when { branch 'main' }
      steps {
        script {
          new flowSemver().commitSemver(VERSION)
        }
      }
    }


    stage('Docker image builds') {
      parallel {
          stage('Upgrade node docker image') {
            when { branch 'main' }
            steps {
              container('node') {
                script {
                  withCredentials([string(credentialsId: "jenkins-hub-api-token", variable: 'GITHUB_TOKEN')]){
                    withAWS(roleAccount: '479720515435', role: 'jenkins-build') {
                      docker.withRegistry('https://index.docker.io/v1/', 'jenkins-dockerhub') {
                        sh """apk add --no-cache ruby curl aws-cli"""
                        sh """cd node && ./build-node ${VERSION.printable()} 12"""
                        sh """cd node && ./build-node_builder ${VERSION.printable()} 12"""
                        sh """cd node && ./build-node ${VERSION.printable()} 16"""
                        sh """cd node && ./build-node_builder ${VERSION.printable()} 16"""
                        sh """cd node && ./build-node ${VERSION.printable()} 18"""
                        sh """cd node && ./build-node_builder ${VERSION.printable()} 18"""
                      }
                    }
                  }
                }
              }
            }
          }
          stage('Upgrade play docker image') {
            when { branch 'main' }
            steps {
              container('play') {
                script {
                  docker.withRegistry('https://index.docker.io/v1/', 'jenkins-dockerhub') {
                    sh """apk add --no-cache curl ruby"""
                    //sh """cd play && ./build-play ${VERSION.printable()} 13"""
                    sh """cd play && ./build-play ${VERSION.printable()} 13-v2 'linux/amd64,linux/arm64'"""
                    sh """cd play && ./build-play ${VERSION.printable()} 17-v2 'linux/amd64,linux/arm64'"""
                  }
                }
              }
            }
          }
          stage('Upgrade play builder docker image') {
            when { branch 'main' }
            steps {
              container('play-builder') {
                script {
                  sh "apk add --no-cache docker-cli openssh curl git ruby"
                  withCredentials([usernamePassword(credentialsId: 'jenkins-x-github', passwordVariable: 'GIT_PASSWORD', usernameVariable: 'GIT_USERNAME')]){
                    withAWS(roleAccount: '479720515435', role: 'jenkins-build') {
                      docker.withRegistry('https://index.docker.io/v1/', 'jenkins-dockerhub') {
                        sh """mkdir /root/.ssh && chmod 0700 /root/.ssh"""
                        sh """ssh-keyscan -H github.com >> ~/.ssh/known_hosts"""
                        sh """cd play && ./build-play-builder ${VERSION.printable()} 13"""
                      }
                    }
                  }
                }
              }
            }
          }
      }
    }
  }
  post {
      failure {
        withCredentials([string(credentialsId: 'slack-team-foundation-notifications-token', variable: 'slackToken')]) {
          slackSend(
            channel: "#team-foundation-notifications", 
            teamDomain: 'flowio.slack.com',
            baseUrl: 'https://flowio.slack.com/services/hooks/jenkins-ci/',
            token: slackToken,
            color: "#ff0000", 
            message: "Build of base docker images failed. Please see https://jenkins.flo.pub/blue/organizations/jenkins/flowcommerce%2Fdocker/activity for details."
          )
        }
      }
  }
}
