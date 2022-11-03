properties([pipelineTriggers([githubPush()])])
pipeline {
  triggers {
    // Only trigger the cron build if on main branch and 5pm friday
    cron(env.BRANCH_NAME == 'main' ? '0 17 * * 5' : '')
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
        containerTemplate(name: 'docker', image: 'docker:20', resourceRequestCpu: '1', resourceRequestMemory: '2Gi', command: 'cat', ttyEnabled: true),
        containerTemplate(name: 'play-builder', image: 'flowdocker/play_builder:latest-java13', resourceRequestCpu: '1', resourceRequestMemory: '2Gi', command: 'cat', ttyEnabled: true)

      ])
    }
  }

  environment {
    ORG = 'flowcommerce'
    APP_NAME = 'docker'
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
            steps {
              container('docker') {
                script{
                  withCredentials([string(credentialsId: "jenkins-hub-api-token", variable: 'GITHUB_TOKEN')]){
                    withAWS(roleAccount: '479720515435', role: 'jenkins-build') {
                      docker.withRegistry('https://index.docker.io/v1/', 'jenkins-dockerhub') {
                        sh """
                            apk update
                            apk add ruby curl aws-cli
                            cd node
                            ./build-node ${VERSION.printable()} 12
                            ./build-node_builder ${VERSION.printable()} 12
                            ./build-node ${VERSION.printable()} 16
                            ./build-node_builder ${VERSION.printable()} 16
                            ./build-node ${VERSION.printable()} 18
                            ./build-node_builder ${VERSION.printable()} 18
                          """
                      }
                    }
                  }
                }
              }
            }
          }
          stage('Upgrade play docker image') {
            steps {
              container('play-builder') {
                script{
                  sh "apk update && apk add --no-cache docker-cli"
                  withCredentials([usernamePassword(credentialsId: 'jenkins-x-github', passwordVariable: 'GIT_PASSWORD', usernameVariable: 'GIT_USERNAME')]){
                    withAWS(roleAccount: '479720515435', role: 'jenkins-build') {
                      docker.withRegistry('https://index.docker.io/v1/', 'jenkins-dockerhub') {
                        sh """
                            apk update
                            apk add --no-cache openssh
                            apk add curl -yqq git ruby
                            mkdir /root/.ssh && chmod 0700 /root/.ssh 
                            ssh-keyscan -H github.com >> ~/.ssh/known_hosts
                            cd play 
                            ./build-play ${VERSION.printable()} 13 
                            ./build-play-builder ${VERSION.printable()} 13
                        """
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
}
