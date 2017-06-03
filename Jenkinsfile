pipeline {
  agent {
    docker {
      image 'golang'
    }
    
  }
  stages {
    stage('test') {
      steps {
        sh 'go get github.com/kazein09/demo'
      }
    }
  }
}