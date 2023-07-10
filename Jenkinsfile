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


    stage('Docker image builds') {
      parallel {
        stage('Upgrade node docker image') {
          steps {
            script {
              withCredentials([string(credentialsId: "jenkins-hub-api-token", variable: 'GITHUB_TOKEN')]){
                withAWS(roleAccount: '479720515435', role: 'jenkins-build') {
                    sh """curl -O https://cdn.flow.io/util/environment-provider/environment-provider-version.txt"""
                    sh """curl -O https://cdn.flow.io/util/environment-provider/environment-provider.jar"""
                    s3Download(file:'./.npmrc', bucket:'io.flow.infra', path:'npm/flowtech.npmrc')
                    sh """cp node/docker/Dockerfile-12 ./Dockerfile"""
                    sh """
                      export semver=testtag &&
                      /kaniko/executor \
                      --dockerfile=./Dockerfile \
                      --context=`pwd` \
                      --snapshot-mode=redo \
                      --use-new-run \
                      --custom-platform=linux/amd64 \
                      --destination flowdocker/node12:$semver
                    """
                    //sh """
                    //  /kaniko/executor \
                    //  --dockerfile=./Dockerfile \
                    //  --context=`pwd` \
                    //  --snapshot-mode=redo \
                    //  --use-new-run \
                    //  --custom-platform=linux/amd64 \
                    //  --destination flowdocker/node12:latest
                    //"""
                    sh """cp node/docker/Dockerfile-16 ./Dockerfile"""
                    sh """
                      export semver=testtag &&
                      /kaniko/executor \
                      --dockerfile=./Dockerfile \
                      --context=`pwd` \
                      --snapshot-mode=redo \
                      --use-new-run \
                      --custom-platform=linux/amd64 \
                      --destination flowdocker/node16:$semver
                    """
                    //sh """
                    //  /kaniko/executor \
                    //  --dockerfile=./Dockerfile \
                    //  --context=`pwd` \
                    //  --snapshot-mode=redo \
                    //  --use-new-run \
                    //  --custom-platform=linux/amd64 \
                    //  --destination flowdocker/node16:latest
                    //"""
                    sh """cp node/docker/Dockerfile-18 ./Dockerfile"""
                    sh """
                      export semver=testtag &&
                      /kaniko/executor \
                      --dockerfile=./Dockerfile \
                      --context=`pwd` \
                      --snapshot-mode=redo \
                      --use-new-run \
                      --custom-platform=linux/amd64 \
                      --destination flowdocker/node18:$semver
                    """
                    //sh """
                    //  /kaniko/executor \
                    //  --dockerfile=./Dockerfile \
                    //  --context=`pwd` \
                    //  --snapshot-mode=redo \
                    //  --use-new-run \
                    //  --custom-platform=linux/amd64 \
                    //  --destination flowdocker/node18:latest
                    //"""
                }
              }
            }
          }
        }
        // stage('Upgrade play jre docker image') {
        //   when { branch 'main' }
        //   steps {
        //     container('play') {
        //       script {
        //         docker.withRegistry('https://index.docker.io/v1/', 'jenkins-dockerhub') {
        //           sh """apk add --no-cache curl ruby"""
        //           sh """cd play && ./build-play ${VERSION.printable()} 13"""
        //           sh """cd play && ./build-play ${VERSION.printable()} 17"""
        //         }
        //       }
        //     }
        //   }
        // }
        // stage('Upgrade play builder docker image') {
        //   when { branch 'main' }
        //   steps {
        //     container('play-builder') {
        //       script {
        //         semver = VERSION.printable()
        //         SBT_VERSION = '1.8.3'
        //         withCredentials([usernamePassword(credentialsId: 'jenkins-x-github', passwordVariable: 'GIT_PASSWORD', usernameVariable: 'GIT_USERNAME')]){
        //           withAWS(roleAccount: '479720515435', role: 'jenkins-build') {
        //             docker.withRegistry('https://index.docker.io/v1/', 'jenkins-dockerhub') {
        //               db = docker.build("flowdocker/play_builder:$semver-java13", "--network=host --build-arg GIT_PASSWORD=$GIT_PASSWORD --build-arg GIT_USERNAME=$GIT_USERNAME --build-arg SBT_VERSION=$SBT_VERSION -f play/Dockerfile.play_builder-13 .")
        //               db.push()
        //               db.push("latest-java13")
        //               db = docker.build("flowdocker/play_builder:$semver-java17", "--network=host --build-arg GIT_PASSWORD=$GIT_PASSWORD --build-arg GIT_USERNAME=$GIT_USERNAME --build-arg SBT_VERSION=$SBT_VERSION -f play/Dockerfile.play_builder-17 .")
        //               db.push()
        //               db.push("latest-java17")
        //             }
        //           }
        //         }
        //       }
        //     }
        //   }
        // }
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
