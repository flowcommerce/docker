properties([pipelineTriggers([githubPush()])])
pipeline {
  options {
    disableConcurrentBuilds()
    buildDiscarder(logRotator(numToKeepStr: '3'))
    timeout(time: 30, unit: 'MINUTES')
  }

  agent {
    kubernetes {
      label 'worker-docker'
      inheritFrom 'default'

      containerTemplates([
        containerTemplate(name: 'golang', image: 'golang:1.12.5', ttyEnabled: true, command: 'cat')
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
      }
    }

    stage('Upgrade node docker image') {
      steps {
        container('golang') {
          script {
            sh('pwd')
            sh('cd node')
            sh('ls')
          }
        }
      }
    }

    stage('Upgrade play docker image') {
      steps {
        container('golang') {
          script {
            sh('pwd')
          }
        }
      }
    }
  }
}
