properties([pipelineTriggers([githubPush()])])
pipeline {
  triggers {
    // Only trigger the cron build if on main branch and 10 AM Monday
    cron(env.BRANCH_NAME == 'main' ? 'TZ=GMT\n0 10 * * 1' : '')
  }
  options {
    disableConcurrentBuilds()
    buildDiscarder(logRotator(numToKeepStr: '3'))
    timeout(time: 60, unit: 'MINUTES')
  }
  agent {
    kubernetes {
      inheritFrom 'default'
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
/*
    stage('Upgrade node docker image 12') {
      when { branch 'main' }
      agent {
        kubernetes {
          label 'docker-image-12'
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
            semver = VERSION.printable()
            env.NODEVERSION = "12"
            sh """/kaniko/executor -f `pwd`/Dockerfile-node -c `pwd` \
              --snapshot-mode=redo --use-new-run  \
              --build-arg NODE_VERSION=${NODEVERSION} \
              --destination flowdocker/node${NODEVERSION}:$semver \
              --destination flowdocker/node${NODEVERSION}:latest
            """
          }
        }
      }
    }

    stage('Upgrade node docker image 16') {
      when { branch 'main' }
      agent {
        kubernetes {
          label 'docker-image-16'
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
            semver = VERSION.printable()
            env.NODEVERSION = "16"
            sh """/kaniko/executor -f `pwd`/Dockerfile-node -c `pwd` \
              --snapshot-mode=redo --use-new-run  \
              --build-arg NODE_VERSION=${NODEVERSION} \
              --destination flowdocker/node${NODEVERSION}:$semver \
              --destination flowdocker/node${NODEVERSION}:latest
            """
          }
        }
      }
    }

    stage('Upgrade node docker image 18') {
      when { branch 'main' }
      agent {
        kubernetes {
          label 'docker-image-18'
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
            semver = VERSION.printable()
            env.NODEVERSION = "18"
            sh """/kaniko/executor -f `pwd`/Dockerfile-node -c `pwd` \
              --snapshot-mode=redo --use-new-run  \
              --build-arg NODE_VERSION=${NODEVERSION} \
              --destination flowdocker/node${NODEVERSION}:$semver \
              --destination flowdocker/node${NODEVERSION}:latest              
            """
          }
        }
      }
    }

    stage('Upgrade node docker image 20') {
      when { branch 'main' }
      agent {
        kubernetes {
          label 'docker-image-20'
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
            semver = VERSION.printable()
            env.NODEVERSION = "20"
            sh """/kaniko/executor -f `pwd`/Dockerfile-node -c `pwd` \
              --snapshot-mode=redo --use-new-run  \
              --build-arg NODE_VERSION=${NODEVERSION} \
              --destination flowdocker/node${NODEVERSION}:$semver \
              --destination flowdocker/node${NODEVERSION}:latest              
            """
          }
        }
      }
    }

    stage('Upgrade node docker builder image 12') {
      when { branch 'main' }
      agent {
        kubernetes {
          label 'docker-builder-image-12'
          inheritFrom 'kaniko-slim'
        }
      }
      steps {
        script {
          withCredentials([string(credentialsId: "jenkins-hub-api-token", variable: 'GITHUB_TOKEN')]){
            withAWS(roleAccount: '479720515435', role: 'jenkins-build') {
              s3Download(file:'./.npmrc', bucket:'io.flow.infra', path:'npm/flowtech.npmrc')
            }
          }
        }
        container('kaniko') {
          script {
            withCredentials([string(credentialsId: "jenkins-hub-api-token", variable: 'GITHUB_TOKEN')]){
              semver = VERSION.printable()
              env.NODEVERSION = "12"
              sh """/kaniko/executor -f `pwd`/Dockerfile-builder -c `pwd` \
                --snapshot-mode=redo --use-new-run  \
                --build-arg NODE_VERSION=${NODEVERSION} \
                --build-arg GITHUB_TOKEN=$GITHUB_TOKEN \
                --destination flowdocker/node${NODEVERSION}_builder:$semver \
                --destination flowdocker/node${NODEVERSION}_builder:latest
              """
            }
          }
        }
      }
    }

    stage('Upgrade node docker builder image 16') {
      when { branch 'main' }
      agent {
        kubernetes {
          label 'docker-builder-image-16'
          inheritFrom 'kaniko-slim'
        }
      }
      steps {
        script {
          withCredentials([string(credentialsId: "jenkins-hub-api-token", variable: 'GITHUB_TOKEN')]){
            withAWS(roleAccount: '479720515435', role: 'jenkins-build') {
              s3Download(file:'./.npmrc', bucket:'io.flow.infra', path:'npm/flowtech.npmrc')
            }
          }
        }
        container('kaniko') {
          script {
            withCredentials([string(credentialsId: "jenkins-hub-api-token", variable: 'GITHUB_TOKEN')]){
              semver = VERSION.printable()
              env.NODEVERSION = "16"
              sh """/kaniko/executor -f `pwd`/Dockerfile-builder -c `pwd` \
                --snapshot-mode=redo --use-new-run  \
                --build-arg NODE_VERSION=${NODEVERSION} \
                --build-arg GITHUB_TOKEN=$GITHUB_TOKEN \
                --destination flowdocker/node${NODEVERSION}_builder:$semver \
                --destination flowdocker/node${NODEVERSION}_builder:latest
              """
            }
          }
        }
      }
    }

    stage('Upgrade node docker builder image 18') {
      when { branch 'main' }
      agent {
        kubernetes {
          label 'docker-builder-image-18'
          inheritFrom 'kaniko-slim'
        }
      }
      steps {
        script {
          withCredentials([string(credentialsId: "jenkins-hub-api-token", variable: 'GITHUB_TOKEN')]){
            withAWS(roleAccount: '479720515435', role: 'jenkins-build') {
              s3Download(file:'./.npmrc', bucket:'io.flow.infra', path:'npm/flowtech.npmrc')
            }
          }
        }
        container('kaniko') {
          script {
            withCredentials([string(credentialsId: "jenkins-hub-api-token", variable: 'GITHUB_TOKEN')]){
              semver = VERSION.printable()
              env.NODEVERSION = "18"
              sh """/kaniko/executor -f `pwd`/Dockerfile-builder -c `pwd` \
                --snapshot-mode=redo --use-new-run  \
                --build-arg NODE_VERSION=${NODEVERSION} \
                --build-arg GITHUB_TOKEN=$GITHUB_TOKEN \
                --destination flowdocker/node${NODEVERSION}_builder:$semver \
                --destination flowdocker/node${NODEVERSION}_builder:latest
              """
            }
          }
        }
      }
    }

    stage('Upgrade node docker builder image 20') {
      when { branch 'main' }
      agent {
        kubernetes {
          label 'docker-builder-image-20'
          inheritFrom 'kaniko-slim'
        }
      }
      steps {
        script {
          withCredentials([string(credentialsId: "jenkins-hub-api-token", variable: 'GITHUB_TOKEN')]){
            withAWS(roleAccount: '479720515435', role: 'jenkins-build') {
              s3Download(file:'./.npmrc', bucket:'io.flow.infra', path:'npm/flowtech.npmrc')
            }
          }
        }
        container('kaniko') {
          script {
            withCredentials([string(credentialsId: "jenkins-hub-api-token", variable: 'GITHUB_TOKEN')]){
              semver = VERSION.printable()
              env.NODEVERSION = "20"
              sh """/kaniko/executor -f `pwd`/Dockerfile-builder -c `pwd` \
                --snapshot-mode=redo --use-new-run  \
                --build-arg NODE_VERSION=${NODEVERSION} \
                --build-arg GITHUB_TOKEN=$GITHUB_TOKEN \
                --destination flowdocker/node${NODEVERSION}_builder:$semver \
                --destination flowdocker/node${NODEVERSION}_builder:latest
              """
            }
          }
        }
      }
    }

    stage('Upgrade docker play java 17') {
      when { branch 'main' }
      agent {
        kubernetes {
          label 'docker-play-17'
          inheritFrom 'kaniko-slim'
        }
      }
      steps {
        script {
          withAWS(roleAccount: '479720515435', role: 'jenkins-build') {
            sh """curl -O https://cdn.flow.io/util/environment-provider/environment-provider-version.txt"""
            sh """curl -O https://cdn.flow.io/util/environment-provider/environment-provider.jar"""
          }
        }
        container('kaniko') {
          script {
            semver = VERSION.printable()
            env.JAVAVERSION = "17"
            sh """/kaniko/executor -f `pwd`/Dockerfile-play-${JAVAVERSION} -c `pwd` \
              --snapshot-mode=redo --use-new-run  \
              --destination flowdocker/play:$semver-java${JAVAVERSION} \
              --destination flowdocker/play:latest-java${JAVAVERSION}
            """
          }
        }
      }
    }

    stage('Upgrade docker play builder java 17') {
      when { branch 'main' }
      agent {
        kubernetes {
          label 'docker-play-builder-17'
          inheritFrom 'kaniko-slim'
        }
      }
      steps {
        container('kaniko') {
          script {
            withCredentials([usernamePassword(credentialsId: 'jenkins-x-github', passwordVariable: 'GIT_PASSWORD', usernameVariable: 'GIT_USERNAME')]){
              semver = VERSION.printable()
              env.JAVAVERSION = "17"
              env.SBT_VERSION = "1.9.9"
              sh """/kaniko/executor -f `pwd`/Dockerfile-play-builder-${JAVAVERSION} -c `pwd` \
                --snapshot-mode=redo --use-new-run  \
                --build-arg SBT_VERSION=${SBT_VERSION} \
                --build-arg GIT_PASSWORD=$GIT_PASSWORD \
                --build-arg GIT_USERNAME=$GIT_USERNAME \
                --destination flowdocker/play_builder:$semver-java${JAVAVERSION} \
                --destination flowdocker/play_builder:latest-java${JAVAVERSION}
              """
            }
          }
        }
      }
    }

    stage('Upgrade docker play builder java 17 - Ubuntu Jammy') {
      when { branch 'main' }
      agent {
        kubernetes {
          label 'docker-play-builder-17-jammy'
          inheritFrom 'kaniko-slim'
        }
      }
      steps {
        container('kaniko') {
          script {
            withCredentials([usernamePassword(credentialsId: 'jenkins-x-github', passwordVariable: 'GIT_PASSWORD', usernameVariable: 'GIT_USERNAME'), string(credentialsId: 'jenkins-hub-api-token', variable: 'GITHUB_TOKEN'), string(credentialsId: 'jenkins-apibuilder-token', variable: 'APIBUILDER_TOKEN') ]) {
              semver = VERSION.printable()
              env.JAVAVERSION = "17"
              env.SBT_VERSION = "1.9.9"
              sh """/kaniko/executor -f `pwd`/Dockerfile-play-builder-${JAVAVERSION}-jammy -c `pwd` \
                --snapshot-mode=redo --use-new-run  \
                --build-arg SBT_VERSION=${SBT_VERSION} \
                --build-arg GIT_PASSWORD=$GIT_PASSWORD \
                --build-arg GIT_USERNAME=$GIT_USERNAME \
                --build-arg GITHUB_TOKEN=$GITHUB_TOKEN \
                --build-arg APIBUILDER_TOKEN=$APIBUILDER_TOKEN \
                --destination flowdocker/play_builder:$semver-java${JAVAVERSION}-jammy \
                --destination flowdocker/play_builder:latest-java${JAVAVERSION}-jammy
              """
            }
          }
        }
      }
    }
*/
    stage('Upgrade docker play builder java 17 arm64 - Ubuntu Jammy') {
      agent {
        kubernetes {
          label 'docker-play-builder-17-jammy-arm64'
          inheritFrom 'default'
          containerTemplates([
              containerTemplate(name: 'kaniko', image: 'gcr.io/kaniko-project/executor:debug', alwaysPullImage: true)
          ])
        }
      }
      steps {
        container('kaniko') {
          script {
            withCredentials([usernamePassword(credentialsId: 'jenkins-x-github', passwordVariable: 'GIT_PASSWORD', usernameVariable: 'GIT_USERNAME'), string(credentialsId: 'jenkins-hub-api-token', variable: 'GITHUB_TOKEN'), string(credentialsId: 'jenkins-apibuilder-token', variable: 'APIBUILDER_TOKEN') ]) {
              semver = VERSION.printable()
              env.JAVAVERSION = "17"
              env.SBT_VERSION = "1.9.9"
              sh """sleep 900; /kaniko/executor -f `pwd`/Dockerfile-play-builder-${JAVAVERSION}-jammy-arm64 -c `pwd` \
                --snapshot-mode=redo --use-new-run  \
                --build-arg SBT_VERSION=${SBT_VERSION} \
                --build-arg GIT_PASSWORD=$GIT_PASSWORD \
                --build-arg GIT_USERNAME=$GIT_USERNAME \
                --build-arg GITHUB_TOKEN=$GITHUB_TOKEN \
                --build-arg APIBUILDER_TOKEN=$APIBUILDER_TOKEN \
                --destination flowdocker/play_builder:$semver-java${JAVAVERSION}-jammy-arm64 \
                --destination flowdocker/play_builder:latest-java${JAVAVERSION}-jammy-arm64 \
                --verbosity debug;
                sleep 900
              """
            }
          }
        }
      }
    }
  }

/*  post {
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
*/  
}