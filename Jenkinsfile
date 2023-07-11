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
      inheritFrom 'kaniko-slim'
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

    stage('Download required files') {
      steps {
        script {
          withCredentials([string(credentialsId: "jenkins-hub-api-token", variable: 'GITHUB_TOKEN')]){
            withAWS(roleAccount: '479720515435', role: 'jenkins-build') {
              sh """curl -O https://cdn.flow.io/util/environment-provider/environment-provider-version.txt"""
              sh """curl -O https://cdn.flow.io/util/environment-provider/environment-provider.jar"""
              s3Download(file:'./.npmrc', bucket:'io.flow.infra', path:'npm/flowtech.npmrc')
            }
          }
        }
      }
    }

    //stage('Docker image builds') {
    //  parallel {
    stage('Upgrade node docker image 12') {
      steps {
        container('kaniko') {
          script {
            sh """cp node/dockerfiles/Dockerfile-12 ./Dockerfile-12 \
              && /kaniko/executor -f `pwd`/Dockerfile-12 -c `pwd` \
              --snapshot-mode=redo --use-new-run  \
              --destination flowdocker/node12:testag
            """
            //sh """cp node/dockerfiles/Dockerfile-12 ./Dockerfile \
            //  && /kaniko/executor -f `pwd`/Dockerfile -c `pwd` \
            //  --snapshot-mode=redo --use-new-run  \
            //  --destination flowdocker/node12:latest
            //"""
          }
        }
      }
    }
    stage('Upgrade node docker image 16') {
      steps {
        container('kaniko') {
          script {
            sh """cp node/dockerfiles/Dockerfile-16 ./Dockerfile-16 \
              && /kaniko/executor -f `pwd`/Dockerfile-16 -c `pwd` \
              --snapshot-mode=redo --use-new-run  \
              --destination flowdocker/node12:testag
            """
            //sh """cp node/dockerfiles/Dockerfile-16 ./Dockerfile \
            //  && /kaniko/executor -f `pwd`/Dockerfile-16 -c `pwd` \
            //  --snapshot-mode=redo --use-new-run  \
            //  --destination flowdocker/node12:latest
            //"""
          }
        }
      }
    }
    stage('Upgrade node docker image 18') {
      steps {
        container('kaniko') {
          script {
            sh """cp node/dockerfiles/Dockerfile-16 ./Dockerfile-18 \
              && /kaniko/executor -f `pwd`/Dockerfile-18 -c `pwd` \
              --snapshot-mode=redo --use-new-run  \
              --destination flowdocker/node12:testag
            """
            //sh """cp node/dockerfiles/Dockerfile-18 ./Dockerfile \
            //  && /kaniko/executor -f `pwd`/Dockerfile-16 -c `pwd` \
            //  --snapshot-mode=redo --use-new-run  \
            //  --destination flowdocker/node12:latest
            //"""
          }
        }
      }
    }
//      }
//    }

    stage('Docker image node builds') {
      parallel {
        stage('Upgrade node-builder docker image 12') {
          steps {
            container('kaniko') {
              script {
                withCredentials([string(credentialsId: "jenkins-hub-api-token", variable: 'GITHUB_TOKEN')]){
                  withAWS(roleAccount: '479720515435', role: 'jenkins-build') {
                    sh """
                      cp node/dockerfiles/Dockerfile-builder-12 ./Dockerfile-builder-12 \
                      && /kaniko/executor -f `pwd`/Dockerfile-builder-12 -c `pwd` \
                      --snapshot-mode=redo --use-new-run  \
                      --build-arg GITHUB_TOKEN=$GITHUB_TOKEN \
                      --destination flowdocker/node12_builder:testag
                    """
                    //sh """cp node/dockerfiles/Dockerfile-12 ./Dockerfile \
                    //  && /kaniko/executor -f `pwd`/Dockerfile -c `pwd` \
                    //  --snapshot-mode=redo --use-new-run  \
                    //  --destination flowdocker/node12:latest
                    //"""
                  }
                }
              }
            }
          }
        }
      }
    }
  }
  //post {
  //  failure {
  //    withCredentials([string(credentialsId: 'slack-team-foundation-notifications-token', variable: 'slackToken')]) {
  //      slackSend(
  //        channel: "#team-foundation-notifications",
  //        teamDomain: 'flowio.slack.com',
  //        baseUrl: 'https://flowio.slack.com/services/hooks/jenkins-ci/',
  //        token: slackToken,
  //        color: "#ff0000",
  //        message: "Build of base docker images failed. Please see https://jenkins.flo.pub/blue/organizations/jenkins/flowcommerce%2Fdocker/activity for details."
  //      )
  //    }
  //  }
  //}
}
