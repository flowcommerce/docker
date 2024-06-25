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

    // Multi archtecture node 12  docker  image build in parallel, manifest-tool push multi platform container image reference
    stage('Multi arch Upgrade node docker image 12') {
      parallel {
        stage('Upgrade node docker image 12 amd64') {
          // when { branch 'main' }
          agent {
            kubernetes {
              label 'docker-image-12'
              inheritFrom 'kaniko-slim'
            }
          }
          steps {
            script {
              withCredentials([string(credentialsId: "jenkins-hub-api-token", variable: 'GITHUB_TOKEN')]){
                withAWS(roleAccount: '479720515435', role: 'jenkins-build') {
                  sh """curl -O https://cdn.flow.io/util/environment-provider/environment-provider-version.txt"""
                  sh """curl -O https://cdn.flow.io/util/environment-provider/environment-provider.jar"""
                  s3Download(file:'./.npmrc', bucket:'io.flow.infra', path:'npm/flowtech.npmrc')
                }
              }
            }
            container('kaniko') {
              script {
                semver = VERSION.printable()
                env.NODEVERSION = "12"
                sh """/kaniko/executor -f `pwd`/Dockerfile-node -c `pwd` \
                  --snapshot-mode=redo --use-new-run  \
                  --build-arg NODE_VERSION=${NODEVERSION} \
                  --destination flowdocker/node${NODEVERSION}-amd64:$semver \
                  --destination flowdocker/node${NODEVERSION}-amd64:latest
                """
              }
            }
          }
        }
        stage('Upgrade node docker image 12 arm64') {
          // when { branch 'main' }
          agent {
            kubernetes {
              label 'docker-image-12-arm64'
              inheritFrom 'kaniko-slim-arm64'
            }
          }
          steps {
            script {
              withCredentials([string(credentialsId: "jenkins-hub-api-token", variable: 'GITHUB_TOKEN')]){
                withAWS(roleAccount: '479720515435', role: 'jenkins-build') {
                  sh """curl -O https://cdn.flow.io/util/environment-provider/environment-provider-version.txt"""
                  sh """curl -O https://cdn.flow.io/util/environment-provider/environment-provider.jar"""
                  s3Download(file:'./.npmrc', bucket:'io.flow.infra', path:'npm/flowtech.npmrc')
                }
              }
            }
            container('kaniko') {
              script {
                semver = VERSION.printable()
                env.NODEVERSION = "12"
                sh """/kaniko/executor -f `pwd`/Dockerfile-node -c `pwd` \
                  --snapshot-mode=redo --use-new-run  \
                  --build-arg NODE_VERSION=${NODEVERSION} \
                  --destination flowdocker/node${NODEVERSION}-arm64:$semver \
                  --destination flowdocker/node${NODEVERSION}-arm64:latest
                """
              }
            }
          }
        }
      }
    }
    
    stage('manifest tool step for Node 12 docker images') {
      //when {branch 'main'}
      agent {
        kubernetes {
          label 'manifest-tool-play-images'
          inheritFrom 'kaniko-slim-arm64'
        }
      }
      steps {
        container('kaniko') {
          script {
            semver = VERSION.printable()
            env.NODEVERSION = "12"
            sh """
              wget https://github.com/estesp/manifest-tool/releases/download/v2.0.8/binaries-manifest-tool-2.0.8.tar.gz
              gunzip binaries-manifest-tool-2.0.8.tar.gz
              tar -xvf binaries-manifest-tool-2.0.8.tar
              mv manifest-tool-linux-arm64 manifest-tool
              chmod +x manifest-tool
              ./manifest-tool push from-args --platforms linux/amd64,linux/arm64 --template flowdocker/node${NODEVERSION}-ARCH:$semver --target flowdocker/node${NODEVERSION}:$semver
              ./manifest-tool push from-args --platforms linux/amd64,linux/arm64 --template flowdocker/node${NODEVERSION}-ARCH:latest --target flowdocker/node${NODEVERSION}:latest
              """
          }
        }
      }
    } 
    
    // Multi archtecture node 16  docker  image build in parallel, manifest-tool push multi platform container image reference
    stage('Multi arch Upgrade node docker image 16') {
      parallel {
        stage('Upgrade node docker image 16 amd64') {
          // when { branch 'main' }
          agent {
            kubernetes {
              label 'docker-image-16'
              inheritFrom 'kaniko-slim'
            }
          }
          steps {
            script {
              withCredentials([string(credentialsId: "jenkins-hub-api-token", variable: 'GITHUB_TOKEN')]){
                withAWS(roleAccount: '479720515435', role: 'jenkins-build') {
                  sh """curl -O https://cdn.flow.io/util/environment-provider/environment-provider-version.txt"""
                  sh """curl -O https://cdn.flow.io/util/environment-provider/environment-provider.jar"""
                  s3Download(file:'./.npmrc', bucket:'io.flow.infra', path:'npm/flowtech.npmrc')
                }
              }
            }
            container('kaniko') {
              script {
                semver = VERSION.printable()
                env.NODEVERSION = "16"
                sh """/kaniko/executor -f `pwd`/Dockerfile-node -c `pwd` \
                  --snapshot-mode=redo --use-new-run  \
                  --build-arg NODE_VERSION=${NODEVERSION} \
                  --destination flowdocker/node${NODEVERSION}-amd64:$semver \
                  --destination flowdocker/node${NODEVERSION}-amd64:latest
                """
              }
            }
          }
        }
        stage('Upgrade node docker image 16 arm64') {
          // when { branch 'main' }
          agent {
            kubernetes {
              label 'docker-image-16-arm64'
              inheritFrom 'kaniko-slim-arm64'
            }
          }
          steps {
            script {
              withCredentials([string(credentialsId: "jenkins-hub-api-token", variable: 'GITHUB_TOKEN')]){
                withAWS(roleAccount: '479720515435', role: 'jenkins-build') {
                  sh """curl -O https://cdn.flow.io/util/environment-provider/environment-provider-version.txt"""
                  sh """curl -O https://cdn.flow.io/util/environment-provider/environment-provider.jar"""
                  s3Download(file:'./.npmrc', bucket:'io.flow.infra', path:'npm/flowtech.npmrc')
                }
              }
            }
            container('kaniko') {
              script {
                semver = VERSION.printable()
                env.NODEVERSION = "16"
                sh """/kaniko/executor -f `pwd`/Dockerfile-node -c `pwd` \
                  --snapshot-mode=redo --use-new-run  \
                  --build-arg NODE_VERSION=${NODEVERSION} \
                  --destination flowdocker/node${NODEVERSION}-arm64:$semver \
                  --destination flowdocker/node${NODEVERSION}-arm64:latest
                """
              }
            }
          }
        }
      }
    }
    
    stage('manifest tool step for Node 16 docker images') {
      //when {branch 'main'}
      agent {
        kubernetes {
          label 'manifest-tool-play-images'
          inheritFrom 'kaniko-slim-arm64'
        }
      }
      steps {
        container('kaniko') {
          script {
            semver = VERSION.printable()
            env.NODEVERSION = "16"
            sh """
              wget https://github.com/estesp/manifest-tool/releases/download/v2.0.8/binaries-manifest-tool-2.0.8.tar.gz
              gunzip binaries-manifest-tool-2.0.8.tar.gz
              tar -xvf binaries-manifest-tool-2.0.8.tar
              mv manifest-tool-linux-arm64 manifest-tool
              chmod +x manifest-tool
              ./manifest-tool push from-args --platforms linux/amd64,linux/arm64 --template flowdocker/node${NODEVERSION}-ARCH:$semver --target flowdocker/node${NODEVERSION}:$semver
              ./manifest-tool push from-args --platforms linux/amd64,linux/arm64 --template flowdocker/node${NODEVERSION}-ARCH:latest --target flowdocker/node${NODEVERSION}:latest
              """
          }
        }
      }
    } 

    // Multi archtecture node 18  docker  image build in parallel, manifest-tool push multi platform container image reference
    stage('Multi arch Upgrade node docker image 18') {
      parallel {
        stage('Upgrade node docker image 18 amd64') {
          // when { branch 'main' }
          agent {
            kubernetes {
              label 'docker-image-18'
              inheritFrom 'kaniko-slim'
            }
          }
          steps {
            script {
              withCredentials([string(credentialsId: "jenkins-hub-api-token", variable: 'GITHUB_TOKEN')]){
                withAWS(roleAccount: '479720515435', role: 'jenkins-build') {
                  sh """curl -O https://cdn.flow.io/util/environment-provider/environment-provider-version.txt"""
                  sh """curl -O https://cdn.flow.io/util/environment-provider/environment-provider.jar"""
                  s3Download(file:'./.npmrc', bucket:'io.flow.infra', path:'npm/flowtech.npmrc')
                }
              }
            }
            container('kaniko') {
              script {
                semver = VERSION.printable()
                env.NODEVERSION = "18"
                sh """/kaniko/executor -f `pwd`/Dockerfile-node -c `pwd` \
                  --snapshot-mode=redo --use-new-run  \
                  --build-arg NODE_VERSION=${NODEVERSION} \
                  --destination flowdocker/node${NODEVERSION}-amd64:$semver \
                  --destination flowdocker/node${NODEVERSION}-amd64:latest              
                """
              }
            }
          }
        }
        stage('Upgrade node docker image 18 arm64') {
          // when { branch 'main' }
          agent {
            kubernetes {
              label 'docker-image-18-arm64'
              inheritFrom 'kaniko-slim-arm64'
            }
          }
          steps {
            script {
              withCredentials([string(credentialsId: "jenkins-hub-api-token", variable: 'GITHUB_TOKEN')]){
                withAWS(roleAccount: '479720515435', role: 'jenkins-build') {
                  sh """curl -O https://cdn.flow.io/util/environment-provider/environment-provider-version.txt"""
                  sh """curl -O https://cdn.flow.io/util/environment-provider/environment-provider.jar"""
                  s3Download(file:'./.npmrc', bucket:'io.flow.infra', path:'npm/flowtech.npmrc')
                }
              }
            }
            container('kaniko') {
              script {
                semver = VERSION.printable()
                env.NODEVERSION = "18"
                sh """/kaniko/executor -f `pwd`/Dockerfile-node -c `pwd` \
                  --snapshot-mode=redo --use-new-run  \
                  --build-arg NODE_VERSION=${NODEVERSION} \
                  --destination flowdocker/node${NODEVERSION}-arm64:$semver \
                  --destination flowdocker/node${NODEVERSION}-arm64:latest              
                """
              }
            }
          }
        }
      }
    }
    
    stage('manifest tool step for Node 18 docker images') {
      //when {branch 'main'}
      agent {
        kubernetes {
          label 'manifest-tool-play-images'
          inheritFrom 'kaniko-slim-arm64'
        }
      }
      steps {
        container('kaniko') {
          script {
            semver = VERSION.printable()
            env.NODEVERSION = "18"
            sh """
              wget https://github.com/estesp/manifest-tool/releases/download/v2.0.8/binaries-manifest-tool-2.0.8.tar.gz
              gunzip binaries-manifest-tool-2.0.8.tar.gz
              tar -xvf binaries-manifest-tool-2.0.8.tar
              mv manifest-tool-linux-arm64 manifest-tool
              chmod +x manifest-tool
              ./manifest-tool push from-args --platforms linux/amd64,linux/arm64 --template flowdocker/node${NODEVERSION}-ARCH:$semver --target flowdocker/node${NODEVERSION}:$semver
              ./manifest-tool push from-args --platforms linux/amd64,linux/arm64 --template flowdocker/node${NODEVERSION}-ARCH:latest --target flowdocker/node${NODEVERSION}:latest
              """
          }
        }
      }
    } 

    // Multi archtecture node 20  docker  image build in parallel, manifest-tool push multi platform container image reference
    stage('Multi arch Upgrade node docker image 20') {
      parallel {
        stage('Upgrade node docker image 20 amd64') {
          // when { branch 'main' }
          agent {
            kubernetes {
              label 'docker-image-20'
              inheritFrom 'kaniko-slim'
            }
          }
          steps {
            script {
              withCredentials([string(credentialsId: "jenkins-hub-api-token", variable: 'GITHUB_TOKEN')]){
                withAWS(roleAccount: '479720515435', role: 'jenkins-build') {
                  sh """curl -O https://cdn.flow.io/util/environment-provider/environment-provider-version.txt"""
                  sh """curl -O https://cdn.flow.io/util/environment-provider/environment-provider.jar"""
                  s3Download(file:'./.npmrc', bucket:'io.flow.infra', path:'npm/flowtech.npmrc')
                }
              }
            }
            container('kaniko') {
              script {
                semver = VERSION.printable()
                env.NODEVERSION = "20"
                sh """/kaniko/executor -f `pwd`/Dockerfile-node -c `pwd` \
                  --snapshot-mode=redo --use-new-run  \
                  --build-arg NODE_VERSION=${NODEVERSION} \
                  --destination flowdocker/node${NODEVERSION}-amd64:$semver \
                  --destination flowdocker/node${NODEVERSION}-amd64:latest              
                """
              }
            }
          }
        }
        stage('Upgrade node docker image 20 arm64') {
          //when { branch 'main' }
          agent {
            kubernetes {
              label 'docker-image-20-arm64'
              inheritFrom 'kaniko-slim-arm64'
            }
          }
          steps {
            script {
              withCredentials([string(credentialsId: "jenkins-hub-api-token", variable: 'GITHUB_TOKEN')]){
                withAWS(roleAccount: '479720515435', role: 'jenkins-build') {
                  sh """curl -O https://cdn.flow.io/util/environment-provider/environment-provider-version.txt"""
                  sh """curl -O https://cdn.flow.io/util/environment-provider/environment-provider.jar"""
                  s3Download(file:'./.npmrc', bucket:'io.flow.infra', path:'npm/flowtech.npmrc')
                }
              }
            }
            container('kaniko') {
              script {
                semver = VERSION.printable()
                env.NODEVERSION = "20"
                sh """/kaniko/executor -f `pwd`/Dockerfile-node -c `pwd` \
                  --snapshot-mode=redo --use-new-run  \
                  --build-arg NODE_VERSION=${NODEVERSION} \
                  --destination flowdocker/node${NODEVERSION}-arm64:$semver \
                  --destination flowdocker/node${NODEVERSION}-arm64:latest              
                """
              }
            }
          }
        }
      }
    }
    
    stage('manifest tool step for Node 20 docker images') {
      //when {branch 'main'}
      agent {
        kubernetes {
          label 'manifest-tool-play-images'
          inheritFrom 'kaniko-slim-arm64'
        }
      }
      steps {
        container('kaniko') {
          script {
            semver = VERSION.printable()
            env.NODEVERSION = "20"
            sh """
              wget https://github.com/estesp/manifest-tool/releases/download/v2.0.8/binaries-manifest-tool-2.0.8.tar.gz
              gunzip binaries-manifest-tool-2.0.8.tar.gz
              tar -xvf binaries-manifest-tool-2.0.8.tar
              mv manifest-tool-linux-arm64 manifest-tool
              chmod +x manifest-tool
              ./manifest-tool push from-args --platforms linux/amd64,linux/arm64 --template flowdocker/node${NODEVERSION}-ARCH:$semver --target flowdocker/node${NODEVERSION}:$semver
              ./manifest-tool push from-args --platforms linux/amd64,linux/arm64 --template flowdocker/node${NODEVERSION}-ARCH:latest --target flowdocker/node${NODEVERSION}:latest
              """
          }
        }
      }
    } 

    // Multi archtecture node 12 docker builder  image build in parallel, manifest-tool push multi platform container image reference
    stage('Multi platform Upgrade node docker builder image 12') {
      parallel {
        stage('Upgrade node docker builder image 12 amd64') {
          //when { branch 'main' }
          agent {
            kubernetes {
              label 'docker-builder-image-12'
              inheritFrom 'kaniko-slim'
            }
          }
          steps {
            script {
              withCredentials([string(credentialsId: "jenkins-hub-api-token", variable: 'GITHUB_TOKEN')]){
                withAWS(roleAccount: '479720515435', role: 'jenkins-build') {
                  s3Download(file:'./.npmrc', bucket:'io.flow.infra', path:'npm/flowtech.npmrc')
                }
              }
            }
            container('kaniko') {
              script {
                withCredentials([string(credentialsId: "jenkins-hub-api-token", variable: 'GITHUB_TOKEN')]){
                  semver = VERSION.printable()
                  env.NODEVERSION = "12"
                  sh """/kaniko/executor -f `pwd`/Dockerfile-builder -c `pwd` \
                    --snapshot-mode=redo --use-new-run  \
                    --build-arg NODE_VERSION=${NODEVERSION} \
                    --build-arg GITHUB_TOKEN=$GITHUB_TOKEN \
                    --destination flowdocker/node${NODEVERSION}_builder_amd64:$semver \
                    --destination flowdocker/node${NODEVERSION}_builder_amd64:latest
                  """
                }
              }
            }
          }
        }
        stage('Upgrade node docker builder image 12 arm64') {
          //when { branch 'main' }
          agent {
            kubernetes {
              label 'docker-builder-image-12-arm64'
              inheritFrom 'kaniko-slim-arm64'
            }
          }
          steps {
            script {
              withCredentials([string(credentialsId: "jenkins-hub-api-token", variable: 'GITHUB_TOKEN')]){
                withAWS(roleAccount: '479720515435', role: 'jenkins-build') {
                  s3Download(file:'./.npmrc', bucket:'io.flow.infra', path:'npm/flowtech.npmrc')
                }
              }
            }
            container('kaniko') {
              script {
                withCredentials([string(credentialsId: "jenkins-hub-api-token", variable: 'GITHUB_TOKEN')]){
                  semver = VERSION.printable()
                  env.NODEVERSION = "12"
                  sh """/kaniko/executor -f `pwd`/Dockerfile-builder -c `pwd` \
                    --snapshot-mode=redo --use-new-run  \
                    --build-arg NODE_VERSION=${NODEVERSION} \
                    --build-arg GITHUB_TOKEN=$GITHUB_TOKEN \
                    --destination flowdocker/node${NODEVERSION}_builder_arm64:$semver \
                    --destination flowdocker/node${NODEVERSION}_builder_arm64:latest
                  """
                }
              }
            }
          }
        }
      }
    }
    
    stage('manifest tool step for Node 12 docker builder images') {
      //when {branch 'main'}
      agent {
        kubernetes {
          label 'manifest-tool-play-images'
          inheritFrom 'kaniko-slim-arm64'
        }
      }
      steps {
        container('kaniko') {
          script {
            semver = VERSION.printable()
            env.NODEVERSION = "12"
            sh """
              wget https://github.com/estesp/manifest-tool/releases/download/v2.0.8/binaries-manifest-tool-2.0.8.tar.gz
              gunzip binaries-manifest-tool-2.0.8.tar.gz
              tar -xvf binaries-manifest-tool-2.0.8.tar
              mv manifest-tool-linux-arm64 manifest-tool
              chmod +x manifest-tool
              ./manifest-tool push from-args --platforms linux/amd64,linux/arm64 --template flowdocker/node${NODEVERSION}_builder_ARCH:$semver --target flowdocker/node${NODEVERSION}_builder:$semver
              ./manifest-tool push from-args --platforms linux/amd64,linux/arm64 --template flowdocker/node${NODEVERSION}_builder_ARCH:latest --target flowdocker/node${NODEVERSION}_builder:latest
              """
          }
        }
      }
    } 

    // Multi archtecture node 16 docker builder  image build in parallel, manifest-tool push multi platform container image reference
    stage('Multi platform Upgrade node docker builder image 16') {
      parallel {
        stage('Upgrade node docker builder image 16 amd64') {
          //when { branch 'main' }
          agent {
            kubernetes {
              label 'docker-builder-image-16'
              inheritFrom 'kaniko-slim'
            }
          }
          steps {
            script {
              withCredentials([string(credentialsId: "jenkins-hub-api-token", variable: 'GITHUB_TOKEN')]){
                withAWS(roleAccount: '479720515435', role: 'jenkins-build') {
                  s3Download(file:'./.npmrc', bucket:'io.flow.infra', path:'npm/flowtech.npmrc')
                }
              }
            }
            container('kaniko') {
              script {
                withCredentials([string(credentialsId: "jenkins-hub-api-token", variable: 'GITHUB_TOKEN')]){
                  semver = VERSION.printable()
                  env.NODEVERSION = "16"
                  sh """/kaniko/executor -f `pwd`/Dockerfile-builder -c `pwd` \
                    --snapshot-mode=redo --use-new-run  \
                    --build-arg NODE_VERSION=${NODEVERSION} \
                    --build-arg GITHUB_TOKEN=$GITHUB_TOKEN \
                    --destination flowdocker/node${NODEVERSION}_builder_amd64:$semver \
                    --destination flowdocker/node${NODEVERSION}_builder_amd64:latest
                  """
                }
              }
            }
          }
        }
        stage('Upgrade node docker builder image 16 arm64') {
          //when { branch 'main' }
          agent {
            kubernetes {
              label 'docker-builder-image-16-arm64'
              inheritFrom 'kaniko-slim-arm64'
            }
          }
          steps {
            script {
              withCredentials([string(credentialsId: "jenkins-hub-api-token", variable: 'GITHUB_TOKEN')]){
                withAWS(roleAccount: '479720515435', role: 'jenkins-build') {
                  s3Download(file:'./.npmrc', bucket:'io.flow.infra', path:'npm/flowtech.npmrc')
                }
              }
            }
            container('kaniko') {
              script {
                withCredentials([string(credentialsId: "jenkins-hub-api-token", variable: 'GITHUB_TOKEN')]){
                  semver = VERSION.printable()
                  env.NODEVERSION = "16"
                  sh """/kaniko/executor -f `pwd`/Dockerfile-builder -c `pwd` \
                    --snapshot-mode=redo --use-new-run  \
                    --build-arg NODE_VERSION=${NODEVERSION} \
                    --build-arg GITHUB_TOKEN=$GITHUB_TOKEN \
                    --destination flowdocker/node${NODEVERSION}_builder_arm64:$semver \
                    --destination flowdocker/node${NODEVERSION}_builder_arm64:latest
                  """
                }
              }
            }
          }
        }
      }
    }
    
    stage('manifest tool step for Node 16 docker builder images') {
      //when {branch 'main'}
      agent {
        kubernetes {
          label 'manifest-tool-play-images'
          inheritFrom 'kaniko-slim-arm64'
        }
      }
      steps {
        container('kaniko') {
          script {
            semver = VERSION.printable()
            env.NODEVERSION = "16"
            sh """
              wget https://github.com/estesp/manifest-tool/releases/download/v2.0.8/binaries-manifest-tool-2.0.8.tar.gz
              gunzip binaries-manifest-tool-2.0.8.tar.gz
              tar -xvf binaries-manifest-tool-2.0.8.tar
              mv manifest-tool-linux-arm64 manifest-tool
              chmod +x manifest-tool
              ./manifest-tool push from-args --platforms linux/amd64,linux/arm64 --template flowdocker/node${NODEVERSION}_builder_ARCH:$semver --target flowdocker/node${NODEVERSION}_builder:$semver
              ./manifest-tool push from-args --platforms linux/amd64,linux/arm64 --template flowdocker/node${NODEVERSION}_builder_ARCH:latest --target flowdocker/node${NODEVERSION}_builder:latest
              """
          }
        }
      }
    } 

    // Multi archtecture node 18 docker builder  image build in parallel, manifest-tool push multi platform container image reference
    stage('Multi platform Upgrade node docker builder image 18') {
      parallel {
        stage('Upgrade node docker builder image 18 amd64') {
          //when { branch 'main' }
          agent {
            kubernetes {
              label 'docker-builder-image-18'
              inheritFrom 'kaniko-slim'
            }
          }
          steps {
            script {
              withCredentials([string(credentialsId: "jenkins-hub-api-token", variable: 'GITHUB_TOKEN')]){
                withAWS(roleAccount: '479720515435', role: 'jenkins-build') {
                  s3Download(file:'./.npmrc', bucket:'io.flow.infra', path:'npm/flowtech.npmrc')
                }
              }
            }
            container('kaniko') {
              script {
                withCredentials([string(credentialsId: "jenkins-hub-api-token", variable: 'GITHUB_TOKEN')]){
                  semver = VERSION.printable()
                  env.NODEVERSION = "18"
                  sh """/kaniko/executor -f `pwd`/Dockerfile-builder -c `pwd` \
                    --snapshot-mode=redo --use-new-run  \
                    --build-arg NODE_VERSION=${NODEVERSION} \
                    --build-arg GITHUB_TOKEN=$GITHUB_TOKEN \
                    --destination flowdocker/node${NODEVERSION}_builder_amd64:$semver \
                    --destination flowdocker/node${NODEVERSION}_builder_amd64:latest
                  """
                }
              }
            }
          }
        }
        stage('Upgrade node docker builder image 18 arm64') {
          //when { branch 'main' }
          agent {
            kubernetes {
              label 'docker-builder-image-18-arm64'
              inheritFrom 'kaniko-slim-arm64'
            }
          }
          steps {
            script {
              withCredentials([string(credentialsId: "jenkins-hub-api-token", variable: 'GITHUB_TOKEN')]){
                withAWS(roleAccount: '479720515435', role: 'jenkins-build') {
                  s3Download(file:'./.npmrc', bucket:'io.flow.infra', path:'npm/flowtech.npmrc')
                }
              }
            }
            container('kaniko') {
              script {
                withCredentials([string(credentialsId: "jenkins-hub-api-token", variable: 'GITHUB_TOKEN')]){
                  semver = VERSION.printable()
                  env.NODEVERSION = "18"
                  sh """/kaniko/executor -f `pwd`/Dockerfile-builder -c `pwd` \
                    --snapshot-mode=redo --use-new-run  \
                    --build-arg NODE_VERSION=${NODEVERSION} \
                    --build-arg GITHUB_TOKEN=$GITHUB_TOKEN \
                    --destination flowdocker/node${NODEVERSION}_builder_arm64:$semver \
                    --destination flowdocker/node${NODEVERSION}_builder_arm64:latest
                  """
                }
              }
            }
          }
        }
      }
    }
    
    stage('manifest tool step for Node 18 docker builder images') {
      //when {branch 'main'}
      agent {
        kubernetes {
          label 'manifest-tool-play-images'
          inheritFrom 'kaniko-slim-arm64'
        }
      }
      steps {
        container('kaniko') {
          script {
            semver = VERSION.printable()
            env.NODEVERSION = "18"
            sh """
              wget https://github.com/estesp/manifest-tool/releases/download/v2.0.8/binaries-manifest-tool-2.0.8.tar.gz
              gunzip binaries-manifest-tool-2.0.8.tar.gz
              tar -xvf binaries-manifest-tool-2.0.8.tar
              mv manifest-tool-linux-arm64 manifest-tool
              chmod +x manifest-tool
              ./manifest-tool push from-args --platforms linux/amd64,linux/arm64 --template flowdocker/node${NODEVERSION}_builder_ARCH:$semver --target flowdocker/node${NODEVERSION}_builder:$semver
              ./manifest-tool push from-args --platforms linux/amd64,linux/arm64 --template flowdocker/node${NODEVERSION}_builder_ARCH:latest --target flowdocker/node${NODEVERSION}_builder:latest
              """
          }
        }
      }
    } 

    // Multi archtecture node  docker builder image build in parallel, manifest-tool push multi platform container image reference
    stage('Multi platform Upgrade node docker builder image 20') {
      parallel {
        stage('Upgrade  amd64 node docker builder image 20') {
          //when { branch 'main' }
          agent {
            kubernetes {
              label 'docker-builder-image-20'
              inheritFrom 'kaniko-slim'
            }
          }
          steps {
            script {
              withCredentials([string(credentialsId: "jenkins-hub-api-token", variable: 'GITHUB_TOKEN')]){
                withAWS(roleAccount: '479720515435', role: 'jenkins-build') {
                  s3Download(file:'./.npmrc', bucket:'io.flow.infra', path:'npm/flowtech.npmrc')
                }
              }
            }
            container('kaniko') {
              script {
                withCredentials([string(credentialsId: "jenkins-hub-api-token", variable: 'GITHUB_TOKEN')]){
                  semver = VERSION.printable()
                  env.NODEVERSION = "20"
                  sh """/kaniko/executor -f `pwd`/Dockerfile-builder -c `pwd` \
                    --snapshot-mode=redo --use-new-run  \
                    --build-arg NODE_VERSION=${NODEVERSION} \
                    --build-arg GITHUB_TOKEN=$GITHUB_TOKEN \
                    --destination flowdocker/node${NODEVERSION}_builder_amd64:$semver \
                    --destination flowdocker/node${NODEVERSION}_builder_amd64:latest
                  """
                }
              }
            }
          }
        }
        stage('Upgrade arm64 node docker builder image 20') {
          //when { branch 'main' }
          agent {
            kubernetes {
              label 'docker-builder-image-20-arm64'
              inheritFrom 'kaniko-slim-arm64'
            }
          }
          steps {
            script {
              withCredentials([string(credentialsId: "jenkins-hub-api-token", variable: 'GITHUB_TOKEN')]){
                withAWS(roleAccount: '479720515435', role: 'jenkins-build') {
                  s3Download(file:'./.npmrc', bucket:'io.flow.infra', path:'npm/flowtech.npmrc')
                }
              }
            }
            container('kaniko') {
              script {
                withCredentials([string(credentialsId: "jenkins-hub-api-token", variable: 'GITHUB_TOKEN')]){
                  semver = VERSION.printable()
                  env.NODEVERSION = "20"
                  sh """/kaniko/executor -f `pwd`/Dockerfile-builder -c `pwd` \
                    --snapshot-mode=redo --use-new-run  \
                    --build-arg NODE_VERSION=${NODEVERSION} \
                    --build-arg GITHUB_TOKEN=$GITHUB_TOKEN \
                    --destination flowdocker/node${NODEVERSION}_builder_arm64:$semver \
                    --destination flowdocker/node${NODEVERSION}_builder_arm64:latest
                  """
                }
              }
            }
          }
        }
      }
    }

    stage('manifest tool step for Node 20 docker images') {
      //when {branch 'main'}
      agent {
        kubernetes {
          label 'manifest-tool-play-images'
          inheritFrom 'kaniko-slim-arm64'
        }
      }
      steps {
        container('kaniko') {
          script {
            semver = VERSION.printable()
            env.NODEVERSION = "20"
            sh """
              wget https://github.com/estesp/manifest-tool/releases/download/v2.0.8/binaries-manifest-tool-2.0.8.tar.gz
              gunzip binaries-manifest-tool-2.0.8.tar.gz
              tar -xvf binaries-manifest-tool-2.0.8.tar
              mv manifest-tool-linux-arm64 manifest-tool
              chmod +x manifest-tool
              ./manifest-tool push from-args --platforms linux/amd64,linux/arm64 --template flowdocker/node${NODEVERSION}_builder_ARCH:$semver --target flowdocker/node${NODEVERSION}_builder:$semver
              ./manifest-tool push from-args --platforms linux/amd64,linux/arm64 --template flowdocker/node${NODEVERSION}_builder_ARCH:latest --target flowdocker/node${NODEVERSION}_builder:latest
              """
          }
        }
      }
    } 
    

// Multi archtecture play docker image build in parallel, manifest-tool push multi platform container image reference
    stage('build Multi arch play java 17 docker images') {
      parallel {
        stage('Upgrade amd64 docker play java 17') {
          //when { branch 'main' }
          agent {
            kubernetes {
              label 'docker-play-17'
              inheritFrom 'kaniko-slim'
            }
          }
          steps {
            script {
              withAWS(roleAccount: '479720515435', role: 'jenkins-build') {
                sh """curl -O https://cdn.flow.io/util/environment-provider/environment-provider-version.txt"""
                sh """curl -O https://cdn.flow.io/util/environment-provider/environment-provider.jar"""
              }
            }
            container('kaniko') {
              script {
                semver = VERSION.printable()
                env.JAVAVERSION = "17"
                sh """/kaniko/executor -f `pwd`/Dockerfile-play-${JAVAVERSION} -c `pwd` \
                  --snapshot-mode=redo --use-new-run  \
                  --destination flowdocker/play-amd64:$semver-java${JAVAVERSION} \
                  --destination flowdocker/play-amd64:latest-java${JAVAVERSION}
                """
              }
            }
          }
        }
        stage('Upgrade arm64 docker play java 17') {
          //when { branch 'main' }
          agent {
            kubernetes {
              label 'docker-play-17-arm64'
              inheritFrom 'kaniko-slim-arm64'
            }
          }
          steps {
            script {
              withAWS(roleAccount: '479720515435', role: 'jenkins-build') {
                sh """curl -O https://cdn.flow.io/util/environment-provider/environment-provider-version.txt"""
                sh """curl -O https://cdn.flow.io/util/environment-provider/environment-provider.jar"""
              }
            }
            container('kaniko') {
              script {
                semver = VERSION.printable()
                env.JAVAVERSION = "17"
                sh """/kaniko/executor -f `pwd`/Dockerfile-play-${JAVAVERSION} -c `pwd` \
                  --snapshot-mode=redo --use-new-run  \
                  --destination flowdocker/play-arm64:$semver-java${JAVAVERSION} \
                  --destination flowdocker/play-arm64:latest-java${JAVAVERSION}
                """
              }
            }
          }
        }
      }
    }
    stage('manifest tool step for play docker images') {
      //when {branch 'main'}
      agent {
        kubernetes {
          label 'manifest-tool-play-images'
          inheritFrom 'kaniko-slim-arm64'
        }
      }
      steps {
        container('kaniko') {
          script {
            semver = VERSION.printable()
            env.JAVAVERSION = "17"
            sh """
              wget https://github.com/estesp/manifest-tool/releases/download/v2.0.8/binaries-manifest-tool-2.0.8.tar.gz
              gunzip binaries-manifest-tool-2.0.8.tar.gz
              tar -xvf binaries-manifest-tool-2.0.8.tar
              mv manifest-tool-linux-arm64 manifest-tool
              chmod +x manifest-tool
              ./manifest-tool push from-args --platforms linux/amd64,linux/arm64 --template flowdocker/play-ARCH:$semver-java${JAVAVERSION} --target flowdocker/play:$semver-java${JAVAVERSION}
              ./manifest-tool push from-args --platforms linux/amd64,linux/arm64 --template flowdocker/play-ARCH:latest-java${JAVAVERSION} --target flowdocker/play:latest-java${JAVAVERSION}
              """
          }
        }
      }
    } 

    // eclipse-temurin:17-jdk-alpine doesn't have arm64 image, we don't support multiplatform play_builder image
    // use play_builder:latest-java20-jammy

    stage('Upgrade amd64 docker play builder java 17') {
      // when { branch 'main' }
      agent {
        kubernetes {
          label 'docker-play-builder-17'
          inheritFrom 'kaniko-slim'
        }
      }
      steps {
        container('kaniko') {
          script {
            withCredentials([usernamePassword(credentialsId: 'jenkins-x-github', passwordVariable: 'GIT_PASSWORD', usernameVariable: 'GIT_USERNAME')]){
              semver = VERSION.printable()
              env.JAVAVERSION = "17"
              env.SBT_VERSION = "1.9.9"
              sh """/kaniko/executor -f `pwd`/Dockerfile-play-builder-${JAVAVERSION} -c `pwd` \
                --snapshot-mode=redo --use-new-run  \
                --build-arg SBT_VERSION=${SBT_VERSION} \
                --build-arg GIT_PASSWORD=$GIT_PASSWORD \
                --build-arg GIT_USERNAME=$GIT_USERNAME \
                --destination flowdocker/play_builder:$semver-java${JAVAVERSION} \
                --destination flowdocker/play_builder:latest-java${JAVAVERSION}
              """
            }
          }
        }
      }
    }

    // Multi archtecture play builder jammy docker image build in parallel, manifest-tool push multi platform container image reference
    stage('build Multi arch docker play builder Jammy java 17 docker images') {
      parallel {
        stage('Upgrade docker play builder java 17 - Ubuntu Jammy') {
          // when { branch 'main' }
          agent {
            kubernetes {
              label 'docker-play-builder-17-jammy'
              inheritFrom 'kaniko-slim'
            }
          }
          steps {
            container('kaniko') {
              script {
                withCredentials([usernamePassword(credentialsId: 'jenkins-x-github', passwordVariable: 'GIT_PASSWORD', usernameVariable: 'GIT_USERNAME'), string(credentialsId: 'jenkins-hub-api-token', variable: 'GITHUB_TOKEN'), string(credentialsId: 'jenkins-apibuilder-token', variable: 'APIBUILDER_TOKEN') ]) {
                  semver = VERSION.printable()
                  env.JAVAVERSION = "17"
                  env.SBT_VERSION = "1.9.9"
                  env.ARCH = "amd64"
                  sh """/kaniko/executor -f `pwd`/Dockerfile-play-builder-${JAVAVERSION}-jammy -c `pwd` \
                    --snapshot-mode=redo --use-new-run  \
                    --build-arg SBT_VERSION=${SBT_VERSION} \
                    --build-arg GIT_PASSWORD=$GIT_PASSWORD \
                    --build-arg GIT_USERNAME=$GIT_USERNAME \
                    --build-arg GITHUB_TOKEN=$GITHUB_TOKEN \
                    --build-arg APIBUILDER_TOKEN=$APIBUILDER_TOKEN \
                    --build-arg ARCH=$ARCH \
                    --destination flowdocker/play_builder_${ARCH}:$semver-java${JAVAVERSION}-jammy \
                    --destination flowdocker/play_builder_${ARCH}:latest-java${JAVAVERSION}-jammy
                  """
                }
              }
            }
          }
        }
        stage('Upgrade docker play builder java 17 - Ubuntu Jammy arm64') {
          // when { branch 'main' }
          agent {
            kubernetes {
              label 'docker-play-builder-17-jammy-arm64'
              inheritFrom 'kaniko-slim-arm64'
            }
          }
          steps {
            container('kaniko') {
              script {
                withCredentials([usernamePassword(credentialsId: 'jenkins-x-github', passwordVariable: 'GIT_PASSWORD', usernameVariable: 'GIT_USERNAME'), string(credentialsId: 'jenkins-hub-api-token', variable: 'GITHUB_TOKEN'), string(credentialsId: 'jenkins-apibuilder-token', variable: 'APIBUILDER_TOKEN') ]) {
                  semver = VERSION.printable()
                  env.JAVAVERSION = "17"
                  env.SBT_VERSION = "1.9.9"
                  env.ARCH = "arm64"
                  sh """/kaniko/executor -f `pwd`/Dockerfile-play-builder-${JAVAVERSION}-jammy -c `pwd` \
                    --snapshot-mode=redo --use-new-run  \
                    --build-arg SBT_VERSION=${SBT_VERSION} \
                    --build-arg GIT_PASSWORD=$GIT_PASSWORD \
                    --build-arg GIT_USERNAME=$GIT_USERNAME \
                    --build-arg GITHUB_TOKEN=$GITHUB_TOKEN \
                    --build-arg APIBUILDER_TOKEN=$APIBUILDER_TOKEN \
                    --build-arg ARCH=$ARCH \
                    --destination flowdocker/play_builder_${ARCH}:$semver-java${JAVAVERSION}-jammy \
                    --destination flowdocker/play_builder_${ARCH}:latest-java${JAVAVERSION}-jammy
                  """
                }
              }
            }
          }
        }
      }
    }
    stage('manifest tool step for play builder jammy docker images') {
      //when {branch 'main'}
      agent {
        kubernetes {
          label 'manifest-tool-play-images'
          inheritFrom 'kaniko-slim-arm64'
        }
      }
      steps {
        container('kaniko') {
          script {
            semver = VERSION.printable()
            env.JAVAVERSION = "17"
            sh """
              wget https://github.com/estesp/manifest-tool/releases/download/v2.0.8/binaries-manifest-tool-2.0.8.tar.gz
              gunzip binaries-manifest-tool-2.0.8.tar.gz
              tar -xvf binaries-manifest-tool-2.0.8.tar
              mv manifest-tool-linux-arm64 manifest-tool
              chmod +x manifest-tool
              ./manifest-tool push from-args --platforms linux/amd64,linux/arm64 --template flowdocker/play_builder_ARCH:$semver-java${JAVAVERSION}-jammy --target flowdocker/play_builder:$semver-java${JAVAVERSION}-jammy
              ./manifest-tool push from-args --platforms linux/amd64,linux/arm64 --template flowdocker/play_builder_ARCH:latest-java${JAVAVERSION}-jammy --target flowdocker/play_builder:latest-java${JAVAVERSION}-jammy
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
