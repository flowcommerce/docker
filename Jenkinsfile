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
      //when { branch 'main' }
      steps {
        script {
          new flowSemver().commitSemver(VERSION)
        }
      }
    }

    stage('Upgrade docker play java 17') {
      //when { branch 'main' }
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
              --destination flowdocker/play-amd64:$semver-java${JAVAVERSION} \
              --destination flowdocker/play-amd64:latest-java${JAVAVERSION}
            """
          }
        }
      }
    }

    stage('Upgrade docker play java 17 arm64') {
      //when { branch 'main' }
      agent {
        kubernetes {
          label 'docker-play-17-arm64'
          inheritFrom 'kaniko-slim-arm64'
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
              --destination flowdocker/play-arm64:$semver-java${JAVAVERSION} \
              --destination flowdocker/play-arm64:latest-java${JAVAVERSION}
            """
            sh """
              wget https://github.com/estesp/manifest-tool/releases/download/v2.0.8/binaries-manifest-tool-2.0.8.tar.gz
              gunzip binaries-manifest-tool-2.0.8.tar.gz
              tar -xvf binaries-manifest-tool-2.0.8.tar
              #sleep 300
              mv manifest-tool-linux-arm64 manifest-tool
              chmod +x manifest-tool
              ./manifest-tool push from-args --platforms linux/amd64,linux/arm64 --template flowdocker/play-ARCH:${semver}-java${JAVAVERSION} --target flowdocker/play:${semver}-java${JAVAVERSION}
              ./manifest-tool push from-args --platforms linux/amd64,linux/arm64 --template flowdocker/play-ARCH:latest-java${JAVAVERSION} --target flowdocker/play:latest-java${JAVAVERSION}
              """
          }
        }
      }
    }
  }
  
  // post {
  //   failure {
  //     withCredentials([string(credentialsId: 'slack-team-foundation-notifications-token', variable: 'slackToken')]) {
  //       slackSend(
  //         channel: "#team-foundation-notifications",
  //         teamDomain: 'flowio.slack.com',
  //         baseUrl: 'https://flowio.slack.com/services/hooks/jenkins-ci/',
  //         token: slackToken,
  //         color: "#ff0000",
  //         message: "Build of base docker images failed. Please see https://jenkins.flo.pub/blue/organizations/jenkins/flowcommerce%2Fdocker/activity for details."
  //       )
  //     }
  //   }
  // }
}
