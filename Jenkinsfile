properties([pipelineTriggers([githubPush()])])
pipeline {
  options {
    disableConcurrentBuilds()
    buildDiscarder(logRotator(numToKeepStr: '3'))
    timeout(time: 30, unit: 'MINUTES')
  }
  parameters {
     string(name: 'VERSION', defaultValue: 'VERSION.printable()', description: 'VERSION.printable()')
     string(name: 'VERSION12', defaultValue: '12', description: 'version')
     string(name: 'VERSION13', defaultValue: '13', description: 'version')
     string(name: 'VERSION16', defaultValue: '16', description: 'version')
     string(name: 'VERSION18', defaultValue: '18', description: 'version')
  }

  agent {
    kubernetes {
      label 'worker-docker'
      inheritFrom 'default'

      containerTemplates([
        containerTemplate(name: 'ruby', image: 'ruby', ttyEnabled: true, command: 'cat')
      ])
    }
  }

  environment {
    ORG = 'flowcommerce'
    APP_NAME = 'docker'
    GOPRIVATE='github.com/flowcommerce'
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


    stage('Upgrade node docker image') {
      steps {
        container('ruby') {
         withCredentials([string(credentialsId: "jenkins-hub-api-token", variable: 'GITHUB_TOKEN')]){
           withAWS(roleAccount: '479720515435', role: 'jenkins-build') {
             sh """
                  cd node
                  ./build-node ${params.VERSION} ${params.VERSION12}
                  ./build-node_builder ${params.VERSION} ${params.VERSION12}
                  ./build-node ${params.VERSION} ${params.VERSION16}
                  ./build-node_builder ${params.VERSION} ${params.VERSION16}
                  ./build-node ${params.VERSION} ${params.VERSION18}
                  ./build-node_builder ${params.VERSION} ${params.VERSION18}
               """
              }
             }
           }
         }
       }
    stage('Upgrade play docker image') {
      steps {
        container('ruby') {
          sh "cd play && ./build-play ${VERSION.printable()} ${params.VERSION13} && ./play/build-play-builder ${VERSION.printable()} ${params.VERSION13}"
        }
      }
    }
  }
}
