pipeline {
    agent any
    
    environment {
        DOCKER_REGISTRY_CREDENTIALS = 'devflaviant_dockerhub_credentials'
        IMAGE_NAME = 'name'
        IMAGE_TAG = 'latest-srjgo'
        DOCKER_USERNAME = 'name'
    }
    
    stages {
        stage('Checkout') {
            steps {
                checkout([$class: 'GitSCM', branches: [[name: '*/dev'], [name: '*/staging'], [name: '*/main']], doGenerateSubmoduleConfigurations: false, extensions: [], submoduleCfg: [], userRemoteConfigs: [[url: 'git@github.com:surajNirala/srj-crud-gin.git']]])
            }
        }

    }
}
