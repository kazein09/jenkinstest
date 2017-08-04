pipeline {
  agent {
    node {
      label 'docker'
    }
    
  }
  stages {
    stage('Build') {
      steps {
        sh "docker image ls"
        sh "bash ./build.sh"
        sh "docker image ls"
      }
    }
  }
}
