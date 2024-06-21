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
    stage('manifest tool step for play docker images') {
      //when {branch 'main'}
      steps {
        container('kaniko') {
          script {
            sh """
              sleep 900
              apk add --update curl
              curl -s -L https://github.com/estesp/manifest-tool/releases/download/v2.0.8/binaries-manifest-tool-2.0.8.tar.gz | tar xvz
              mv manifest-tool-linux-amd64 manifest-tool
              chmod +x manifest-tool
              ./manifest-tool push from-args --platforms linux/amd64,linux/arm64 --template flowcommerce/play-ARCH:0.2.83-java17-ARCH --target flowcommerce/play-test:0.2.83-java17-ARCH
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
