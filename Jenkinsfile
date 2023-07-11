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

    stage('Upgrade docker play java 13') {
      agent {
        kubernetes {
          label 'docker-play-13'
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
            env.JAVAVERSION = "13"
            sh """/kaniko/executor -f `pwd`/Dockerfile-play-${JAVAVERSION} -c `pwd` \
              --snapshot-mode=redo --use-new-run  \
              --destination flowdocker/play:testtag-java${JAVAVERSION}
            """
          }
        }
      }
    }

    // There is a destructive action when creating the version, we need to create latest in different agent
    stage('Upgrade docker play java 13 latest') {
      agent {
        kubernetes {
          label 'docker-play-13'
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
            env.JAVAVERSION = "13"
            sh """/kaniko/executor -f `pwd`/Dockerfile-play-${JAVAVERSION} -c `pwd` \
              --snapshot-mode=redo --use-new-run  \
              --destination flowdocker/play:latesttesttag-java${JAVAVERSION}
            """
          }
        }
      }
    }

    stage('Upgrade docker play java 17') {
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
              --destination flowdocker/play:testtag-java${JAVAVERSION}
            """
            sh """/kaniko/executor -f `pwd`/Dockerfile-play-${JAVAVERSION} -c `pwd` \
              --snapshot-mode=redo --use-new-run  \
              --destination flowdocker/play:latesttesttag-java${JAVAVERSION}
            """
          }
        }
      }
    }

    stage('Upgrade docker play builder java 13') {
      agent {
        kubernetes {
          label 'docker-play-builder-13'
          inheritFrom 'kaniko-slim'
        }
      }
      steps {
        container('kaniko') {
          script {
            withCredentials([usernamePassword(credentialsId: 'jenkins-x-github', passwordVariable: 'GIT_PASSWORD', usernameVariable: 'GIT_USERNAME')]){
              semver = VERSION.printable()
              env.JAVAVERSION = "13"
              env.SBT_VERSION = "1.8.3"
              sh """/kaniko/executor -f `pwd`/Dockerfile-play-builder-${JAVAVERSION} -c `pwd` \
                --snapshot-mode=redo --use-new-run  \
                --build-arg SBT_VERSION=${SBT_VERSION} \
                --build-arg GIT_PASSWORD=$GIT_PASSWORD \
                --build-arg GIT_USERNAME=$GIT_USERNAME \
                --destination flowdocker/play_builder:testtag-java${JAVAVERSION}
              """
              sh """/kaniko/executor -f `pwd`/Dockerfile-play-${JAVAVERSION} -c `pwd` \
                --snapshot-mode=redo --use-new-run  \
                --build-arg SBT_VERSION=${SBT_VERSION} \
                --build-arg GIT_PASSWORD=$GIT_PASSWORD \
                --build-arg GIT_USERNAME=$GIT_USERNAME \
                --destination flowdocker/play_builder:latesttesttag-java${JAVAVERSION}
              """
            }
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
              env.SBT_VERSION = "1.8.3"
              sh """/kaniko/executor -f `pwd`/Dockerfile-play-builder-${JAVAVERSION} -c `pwd` \
                --snapshot-mode=redo --use-new-run  \
                --build-arg SBT_VERSION=${SBT_VERSION} \
                --build-arg GIT_PASSWORD=$GIT_PASSWORD \
                --build-arg GIT_USERNAME=$GIT_USERNAME \
                --destination flowdocker/play_builder:testtag-java${JAVAVERSION}
              """
              sh """/kaniko/executor -f `pwd`/Dockerfile-play-${JAVAVERSION} -c `pwd` \
                --snapshot-mode=redo --use-new-run  \
                --build-arg SBT_VERSION=${SBT_VERSION} \
                --build-arg GIT_PASSWORD=$GIT_PASSWORD \
                --build-arg GIT_USERNAME=$GIT_USERNAME \
                --destination flowdocker/play_builder:latesttesttag-java${JAVAVERSION}
              """
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
