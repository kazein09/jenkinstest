pipeline {
  agent {
    docker {
      image 'centos'
    }
    
  }
  stages {
    stage('test') {
      steps {
        sh 'id'
      }
    }
    stage('build') {
      steps {
        parallel(
          "build": {
            sh 'cat file1'
            
          },
          "testbuild": {
            sh 'ls -all'
            
          }
        )
      }
    }
    stage('build 2') {
      steps {
        parallel(
          "build 2": {
            sh 'cat file2'
            
          },
          "testbuild2": {
            sh 'ls -all'
            
          }
        )
      }
    }
  }
}