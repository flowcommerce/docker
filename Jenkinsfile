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
      when { branch 'FDTN-78' } 
      steps {
        container('golang') {
          script {
            sh(script: 'pwd')
            #sh(script: 'cd node')
            #sh(script: 'go run build.go')
          }
        }
      }
    }

    stage('Upgrade play docker image') {
      when { branch 'FDTN-78' } 
      steps {
        container('golang') {
          script {
            sh(script: 'pwd')
            #sh(script: 'cd play')
            #sh(script: 'go run build.go')
          }
        }
      }
    }
  }
}
