properties([pipelineTriggers([githubPush()])])
pipeline {
  options {
    disableConcurrentBuilds()
    buildDiscarder(logRotator(numToKeepStr: '3'))
    timeout(time: 30, unit: 'MINUTES')
  }
  parameters {
     string(name: 'SEM_INFO', defaultValue: 'sem-info tag latest', description: 'SEM-INFO ')
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
             sh "ls"
             sh "./node/build-node ${params.SEM_INFO} ${params.VERSION12}"
             sh "./node/build-node_builder ${params.SEM_INFO} ${params.VERSION12}"
             sh "./node/build-node ${params.SEM_INFO} ${params.VERSION16}"
             sh "./node/build-node_builder ${params.SEM_INFO} ${params.VERSION16}"
             sh "./node/build-node ${params.SEM_INFO} ${params.VERSION18}"
             sh "./node/build-node_builder ${params.SEM_INFO} ${params.VERSION18}"
        }
      }
    }

    stage('Upgrade play docker image') {
      steps {
        container('ruby') {
          sh "ls"
          sh "./play//build-play ${params.SEM_INFO} ${params.VERSION13}"
          sh "./play/build-play-builder ${params.SEM_INFO} ${params.VERSION13}"
          
        }
      }
    }
  }
}
