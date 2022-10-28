properties([pipelineTriggers([githubPush()])])
pipeline {
  options {
    disableConcurrentBuilds()
    buildDiscarder(logRotator(numToKeepStr: '3'))
    timeout(time: 30, unit: 'MINUTES')
  }
  parameters {
     string(name: 'SEM_INFO', defaultValue: 'sem-info tag latest', description: 'SEM-INFO ')
     string(name: 'DB_NAME', defaultValue: 'xyz', description: 'Enter the database that needs to be created')
     string(name: 'VERSION12', defaultValue: '12', description: 'version')
     choice(name: 'RUN', choices: 'Migrate Data', description: 'Data migration')
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
            sh "cd node"
            sh "./node/build-node `sem-info tag latest` 12"
            sh "./build-node_builder `sem-info tag latest` 12"
            sh "./build-node `sem-info tag latest` 16"
            sh "./build-node_builder `sem-info tag latest` 16"
            sh "./build-node `sem-info tag latest` 18"
            sh "./build-node_builder `sem-info tag latest` 18"
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
