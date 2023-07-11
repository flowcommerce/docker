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

//    stage('Upgrade node docker image 12') {
//      when { branch 'main' }
//      agent {
//        kubernetes {
//          label 'docker-image-12'
//          inheritFrom 'kaniko-slim'
//        }
//      }
//      steps {
//        script {
//          withCredentials([string(credentialsId: "jenkins-hub-api-token", variable: 'GITHUB_TOKEN')]){
//            withAWS(roleAccount: '479720515435', role: 'jenkins-build') {
//              sh """curl -O https://cdn.flow.io/util/environment-provider/environment-provider-version.txt"""
//              sh """curl -O https://cdn.flow.io/util/environment-provider/environment-provider.jar"""
//              s3Download(file:'./.npmrc', bucket:'io.flow.infra', path:'npm/flowtech.npmrc')
//            }
//          }
//        }
//        container('kaniko') {
//          script {
//            semver = VERSION.printable()
//            env.NODEVERSION = "12"
//            sh """/kaniko/executor -f `pwd`/Dockerfile -c `pwd` \
//              --snapshot-mode=redo --use-new-run  \
//              --build-arg NODE_VERSION=${NODEVERSION} \
//              --destination flowdocker/node${NODEVERSION}:$semver
//            """
//            sh """/kaniko/executor -f `pwd`/Dockerfile -c `pwd` \
//              --snapshot-mode=redo --use-new-run  \
//              --build-arg NODE_VERSION=${NODEVERSION} \
//              --destination flowdocker/node${NODEVERSION}:latest
//            """
//          }
//        }
//      }
//    }
//
//    stage('Upgrade node docker image 16') {
//      when { branch 'main' }
//      agent {
//        kubernetes {
//          label 'docker-image-16'
//          inheritFrom 'kaniko-slim'
//        }
//      }
//      steps {
//        script {
//          withCredentials([string(credentialsId: "jenkins-hub-api-token", variable: 'GITHUB_TOKEN')]){
//            withAWS(roleAccount: '479720515435', role: 'jenkins-build') {
//              sh """curl -O https://cdn.flow.io/util/environment-provider/environment-provider-version.txt"""
//              sh """curl -O https://cdn.flow.io/util/environment-provider/environment-provider.jar"""
//              s3Download(file:'./.npmrc', bucket:'io.flow.infra', path:'npm/flowtech.npmrc')
//            }
//          }
//        }
//        container('kaniko') {
//          script {
//            semver = VERSION.printable()
//            env.NODEVERSION = "16"
//            sh """/kaniko/executor -f `pwd`/Dockerfile -c `pwd` \
//              --snapshot-mode=redo --use-new-run  \
//              --build-arg NODE_VERSION=${NODEVERSION} \
//              --destination flowdocker/node${NODEVERSION}:$semver
//            """
//            sh """/kaniko/executor -f `pwd`/Dockerfile -c `pwd` \
//              --snapshot-mode=redo --use-new-run  \
//              --build-arg NODE_VERSION=${NODEVERSION} \
//              --destination flowdocker/node${NODEVERSION}:latest
//            """
//          }
//        }
//      }
//    }
//
//    stage('Upgrade node docker image 18') {
//      when { branch 'main' }
//      agent {
//        kubernetes {
//          label 'docker-image-18'
//          inheritFrom 'kaniko-slim'
//        }
//      }
//      steps {
//        script {
//          withCredentials([string(credentialsId: "jenkins-hub-api-token", variable: 'GITHUB_TOKEN')]){
//            withAWS(roleAccount: '479720515435', role: 'jenkins-build') {
//              sh """curl -O https://cdn.flow.io/util/environment-provider/environment-provider-version.txt"""
//              sh """curl -O https://cdn.flow.io/util/environment-provider/environment-provider.jar"""
//              s3Download(file:'./.npmrc', bucket:'io.flow.infra', path:'npm/flowtech.npmrc')
//            }
//          }
//        }
//        container('kaniko') {
//          script {
//            semver = VERSION.printable()
//            env.NODEVERSION = "18"
//            sh """/kaniko/executor -f `pwd`/Dockerfile -c `pwd` \
//              --snapshot-mode=redo --use-new-run  \
//              --build-arg NODE_VERSION=${NODEVERSION} \
//              --destination flowdocker/node${NODEVERSION}:$semver
//            """
//            sh """/kaniko/executor -f `pwd`/Dockerfile -c `pwd` \
//              --snapshot-mode=redo --use-new-run  \
//              --build-arg NODE_VERSION=${NODEVERSION} \
//              --destination flowdocker/node${NODEVERSION}:latest
//            """
//          }
//        }
//      }
//    }
//
//    stage('Upgrade node docker builder image 12') {
//      when { branch 'main' }
//      agent {
//        kubernetes {
//          label 'docker-builder-image-12'
//          inheritFrom 'kaniko-slim'
//        }
//      }
//      steps {
//        script {
//          withCredentials([string(credentialsId: "jenkins-hub-api-token", variable: 'GITHUB_TOKEN')]){
//            withAWS(roleAccount: '479720515435', role: 'jenkins-build') {
//              s3Download(file:'./.npmrc', bucket:'io.flow.infra', path:'npm/flowtech.npmrc')
//            }
//          }
//        }
//        container('kaniko') {
//          script {
//            withCredentials([string(credentialsId: "jenkins-hub-api-token", variable: 'GITHUB_TOKEN')]){
//              semver = VERSION.printable()
//              env.NODEVERSION = "12"
//              sh """/kaniko/executor -f `pwd`/Dockerfile-builder -c `pwd` \
//                --snapshot-mode=redo --use-new-run  \
//                --build-arg NODE_VERSION=${NODEVERSION} \
//                --build-arg GITHUB_TOKEN=$GITHUB_TOKEN \
//                --destination flowdocker/node${NODEVERSION}_builder:$semver
//              """
//              sh """/kaniko/executor -f `pwd`/Dockerfile-builder -c `pwd` \
//                --snapshot-mode=redo --use-new-run  \
//                --build-arg NODE_VERSION=${NODEVERSION} \
//                --build-arg GITHUB_TOKEN=$GITHUB_TOKEN \
//                --destination flowdocker/node${NODEVERSION}_builder:latest
//              """
//            }
//          }
//        }
//      }
//    }
//
//    stage('Upgrade node docker builder image 16') {
//      when { branch 'main' }
//      agent {
//        kubernetes {
//          label 'docker-builder-image-16'
//          inheritFrom 'kaniko-slim'
//        }
//      }
//      steps {
//        script {
//          withCredentials([string(credentialsId: "jenkins-hub-api-token", variable: 'GITHUB_TOKEN')]){
//            withAWS(roleAccount: '479720515435', role: 'jenkins-build') {
//              s3Download(file:'./.npmrc', bucket:'io.flow.infra', path:'npm/flowtech.npmrc')
//            }
//          }
//        }
//        container('kaniko') {
//          script {
//            withCredentials([string(credentialsId: "jenkins-hub-api-token", variable: 'GITHUB_TOKEN')]){
//              semver = VERSION.printable()
//              env.NODEVERSION = "16"
//              sh """/kaniko/executor -f `pwd`/Dockerfile-builder -c `pwd` \
//                --snapshot-mode=redo --use-new-run  \
//                --build-arg NODE_VERSION=${NODEVERSION} \
//                --build-arg GITHUB_TOKEN=$GITHUB_TOKEN \
//                --destination flowdocker/node${NODEVERSION}_builder:$semver
//              """
//              sh """/kaniko/executor -f `pwd`/Dockerfile-builder -c `pwd` \
//                --snapshot-mode=redo --use-new-run  \
//                --build-arg NODE_VERSION=${NODEVERSION} \
//                --build-arg GITHUB_TOKEN=$GITHUB_TOKEN \
//                --destination flowdocker/node${NODEVERSION}_builder:latest
//              """
//            }
//          }
//        }
//      }
//    }
//
//    stage('Upgrade node docker builder image 18') {
//      when { branch 'main' }
//      agent {
//        kubernetes {
//          label 'docker-builder-image-18'
//          inheritFrom 'kaniko-slim'
//        }
//      }
//      steps {
//        script {
//          withCredentials([string(credentialsId: "jenkins-hub-api-token", variable: 'GITHUB_TOKEN')]){
//            withAWS(roleAccount: '479720515435', role: 'jenkins-build') {
//              s3Download(file:'./.npmrc', bucket:'io.flow.infra', path:'npm/flowtech.npmrc')
//            }
//          }
//        }
//        container('kaniko') {
//          script {
//            withCredentials([string(credentialsId: "jenkins-hub-api-token", variable: 'GITHUB_TOKEN')]){
//              semver = VERSION.printable()
//              env.NODEVERSION = "18"
//              sh """/kaniko/executor -f `pwd`/Dockerfile-builder -c `pwd` \
//                --snapshot-mode=redo --use-new-run  \
//                --build-arg NODE_VERSION=${NODEVERSION} \
//                --build-arg GITHUB_TOKEN=$GITHUB_TOKEN \
//                --destination flowdocker/node${NODEVERSION}_builder:$semver
//              """
//              sh """/kaniko/executor -f `pwd`/Dockerfile-builder -c `pwd` \
//                --snapshot-mode=redo --use-new-run  \
//                --build-arg NODE_VERSION=${NODEVERSION} \
//                --build-arg GITHUB_TOKEN=$GITHUB_TOKEN \
//                --destination flowdocker/node${NODEVERSION}_builder:latest
//              """
//            }
//          }
//        }
//      }
//    }

    stage('Upgrade node docker play java 13') {
      agent {
        kubernetes {
          label 'docker-play-13'
          inheritFrom 'kaniko-slim'
        }
      }
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
        container('kaniko') {
          script {
            withCredentials([string(credentialsId: "jenkins-hub-api-token", variable: 'GITHUB_TOKEN')]){
              semver = VERSION.printable()
              env.JAVAVERSION = "13"
              sh """/kaniko/executor -f `pwd`/Dockerfile-play-${JAVAVERSION} -c `pwd` \
                --snapshot-mode=redo --use-new-run  \
                --build-arg NODE_VERSION=${JAVAVERSION} \
                --destination flowdocker/play:testtag-java${JAVAVERSION}
              """
              //sh """/kaniko/executor -f `pwd`/Dockerfile-play-${JAVAVERSION} -c `pwd` \
              //  --snapshot-mode=redo --use-new-run  \
              //  --build-arg NODE_VERSION=${JAVAVERSION} \
              //  --destination flowdocker/play:latest-java${JAVAVERSION}
              //"""
            }
          }
        }
      }
    }

    stage('Upgrade node docker play java 17') {
      agent {
        kubernetes {
          label 'docker-play-17'
          inheritFrom 'kaniko-slim'
        }
      }
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
        container('kaniko') {
          script {
            withCredentials([string(credentialsId: "jenkins-hub-api-token", variable: 'GITHUB_TOKEN')]){
              semver = VERSION.printable()
              env.JAVAVERSION = "17"
              sh """/kaniko/executor -f `pwd`/Dockerfile-play-${JAVAVERSION} -c `pwd` \
                --snapshot-mode=redo --use-new-run  \
                --build-arg NODE_VERSION=${JAVAVERSION} \
                --destination flowdocker/play:testtag-java${JAVAVERSION}
              """
              //sh """/kaniko/executor -f `pwd`/Dockerfile-play-${JAVAVERSION} -c `pwd` \
              //  --snapshot-mode=redo --use-new-run  \
              //  --build-arg NODE_VERSION=${JAVAVERSION} \
              //  --destination flowdocker/play:latest-java${JAVAVERSION}
              //"""
            }
          }
        }
      }
    }

//      }
//    }

//    stage('Docker image node builds') {
//      parallel {
//        stage('Upgrade node-builder docker image 12') {
//          steps {
//            container('kaniko') {
//              script {
//                withCredentials([string(credentialsId: "jenkins-hub-api-token", variable: 'GITHUB_TOKEN')]){
//                  withAWS(roleAccount: '479720515435', role: 'jenkins-build') {
//                    sh """
//                      cp node/dockerfiles/Dockerfile-builder-12 ./Dockerfile-builder-12 \
//                      && /kaniko/executor -f `pwd`/Dockerfile-builder-12 -c `pwd` \
//                      --snapshot-mode=redo --use-new-run  \
//                      --build-arg GITHUB_TOKEN=$GITHUB_TOKEN \
//                      --destination flowdocker/node12_builder:testag
//                    """
//                    //sh """cp node/dockerfiles/Dockerfile-12 ./Dockerfile \
//                    //  && /kaniko/executor -f `pwd`/Dockerfile -c `pwd` \
//                    //  --snapshot-mode=redo --use-new-run  \
//                    //  --destination flowdocker/node12:latest
//                    //"""
//                  }
//                }
//              }
//            }
//          }
//        }
//      }
//    }
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
