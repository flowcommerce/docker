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
        containerTemplate(name: 'ruby', image: 'bitnami/ruby', ttyEnabled: true, command: 'cat')
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
        container('ruby') {
          script {
            sh '''
               "cd node"
               ls
               "./build-node `sem-info tag latest` 12"
               "./build-node_builder `sem-info tag latest` 12"
               "./build-node `sem-info tag latest` 16"
               "./build-node_builder `sem-info tag latest` 16"
               "./build-node `sem-info tag latest` 18"
               "./build-node_builder `sem-info tag latest` 18"
            '''
          }
        }
      }
    }

    stage('Upgrade play docker image') {
      steps {
        container('ruby') {
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
