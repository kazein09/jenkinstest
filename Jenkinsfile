pipeline {
  agent {
    docker {
      image 'golang'
    }
    
  }
  stages {
    stage('test') {
      steps {
        parallel(
          "test": {
            sh 'go build ./main.go'
            
          },
          "error": {
            sh 'ls -all'
            
          }
        )
      }
    }
    stage('build') {
      steps {
        parallel(
          "build": {
            sh '''cat << EOF > file
cd "$HOME"
echo "$PWD" # echo the current path
EOF
ls -all'''
            
          },
          "testbuild": {
            sh 'ls -all'
            
          },
          "new": {
            sh 'exit 0'
            
          }
        )
      }
    }
    stage('build 2') {
      steps {
        parallel(
          "build 2": {
            sh 'exit 1'
            
          },
          "testbuild2": {
            sh 'ls -all'
            
          }
        )
      }
    }
    stage('Test-Nuriel') {
      steps {
        sh 'go test ./main.go'
      }
    }
    stage('curl') {
      steps {
        parallel(
          "curl": {
            sh 'exit 0'
            
          },
          "curl2": {
            sh 'exit 0'
            
          }
        )
      }
    }
  }
}