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
           withCredentials([
             usernamePassword(credentialsId: 'jenkins-x-github', GITHUB_TOKEN: 'GITHUB_TOKEN')
           ]) {
          sh """
               cd node
               ./build-node ${params.SEM_INFO} ${params.VERSION12}
               ./build-node_builder ${params.SEM_INFO} ${params.VERSION12}
               ./build-node ${params.SEM_INFO} ${params.VERSION16}
               ./build-node_builder ${params.SEM_INFO} ${params.VERSION16}
               ./build-node ${params.SEM_INFO} ${params.VERSION18}
               ./build-node_builder ${params.SEM_INFO} ${params.VERSION18}
            """
           }
        }
      }
    }
    stage('Upgrade play docker image') {
      steps {
        container('ruby') {
          sh "cd play && ./build-play ${params.SEM_INFO} ${params.VERSION13} && ./play/build-play-builder ${params.SEM_INFO} ${params.VERSION13}"
        }
      }
    }
  }
}
