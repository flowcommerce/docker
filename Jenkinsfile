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
              env.SBT_VERSION = "1.9.6"
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
  }
}
