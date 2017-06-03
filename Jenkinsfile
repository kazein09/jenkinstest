pipeline {
  agent {
    docker {
      image 'golang'
    }
    
  }
  stages {
    stage('test') {
      steps {
        sh 'echo "passed"'
      }
    }
    stage('build') {
      steps {
        parallel(
          "build": {
            sh '''echo "starting"

go get -d -v github.com/kazein09/demo'''
            
          },
          "testbuild": {
            sh 'echo "lyalya"'
            
          }
        )
      }
    }
  }
}