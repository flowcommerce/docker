properties([pipelineTriggers([githubPush()])])
pipeline {
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
        containerTemplate(name: 'ubuntu', image: 'ubuntu', resourceRequestCpu: '1', resourceRequestMemory: '2Gi', command: 'cat', ttyEnabled: true)

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
            when { branch 'main' }
            steps {
              container('docker') {
                script{
                  withCredentials([usernamePassword(credentialsId: 'jenkins-x-github', passwordVariable: 'GIT_PASSWORD', usernameVariable: 'GIT_USERNAME')]){
                    withAWS(roleAccount: '479720515435', role: 'jenkins-build') {
                      docker.withRegistry('https://index.docker.io/v1/', 'jenkins-dockerhub') {
                        sh """
                            apk update
                            apk add ruby curl aws-cli
                            aws sts get-caller-identity
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
              container('ubuntu') {
                script{
                  sh"apt-get update"
                  sh "apt-get install -y docker.io"
                  withCredentials([usernamePassword(credentialsId: 'jenkins-x-github', passwordVariable: 'GIT_PASSWORD', usernameVariable: 'GIT_USERNAME')]){
                    withAWS(roleAccount: '479720515435', role: 'jenkins-build') {
                      docker.withRegistry('https://index.docker.io/v1/', 'jenkins-dockerhub') {
                        sh """
                            mkdir /root/.ssh && chmod 0700 /root/.ssh 
                            ssh-keyscan -H github.com >> ~/.ssh/known_hosts

                            apt-get update
                            apt-get install apt-transport-https curl gnupg -yqq git ruby
                            echo "deb https://repo.scala-sbt.org/scalasbt/debian all main" | tee /etc/apt/sources.list.d/sbt.list
                            echo "deb https://repo.scala-sbt.org/scalasbt/debian /" | tee /etc/apt/sources.list.d/sbt_old.list
                            curl -sL "https://keyserver.ubuntu.com/pks/lookup?op=get&search=0x2EE0EA64E40A89B84B2DF73499E82A75642AC823" | gpg --no-default-keyring --keyring gnupg-ring:/etc/apt/trusted.gpg.d/scalasbt-release.gpg --import
                            chmod 644 /etc/apt/trusted.gpg.d/scalasbt-release.gpg
                            apt-get update
                            apt-get install sbt -y
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
