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
    GOPRIVATE='github.com/flowcommerce'
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
            sh '''
               cd node
               ls
               git config --global --add url."git@github.com:".insteadOf "https://github.com/"
               go run build.go
            '''
          }
        }
      }
    }

    stage('Upgrade play docker image') {
      steps {
        container('golang') {
          script {
            sh'''
              cd play
              ls
            '''
          }
        }
      }
    }
  }
}
