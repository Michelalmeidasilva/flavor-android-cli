#Requirements


## Args: 

1. BUNDLE_ID
2. APP_FLAVOR
3. BUILD_OUTPUT_TYPE
4. APP_KEY_ALIAS
5. APP_KEY_PASSWORD
6. APP_KEY_STORE_PASSWORD
7. APP_NAME
8. DEEP_LINKING_TAG


## Definir:
1. Como tratar os diferentes locais de .envs
2. Como lidar com erros de utilização do usuário: Rodar a CLI sem estar em paths da aplicação/Informar Args errados.


## Android
1. Duplicar a pasta base em <root>/android/src/base 
2. Modificar o nome da pasta com o ${AppFlavor} 
3. Modificar o app_name de Strings.xml com a ${AppName}
4. Adicionar em base/AndroidManifest.xml a tag de deep-linking
5. Adicionar em .envs do fastlane: BUNDLE_ID, APP_FLAVOR, BUILD_OUTPUT_TYPE, APP_KEY_ALIAS, APP_KEY_PASSWORD. 

## APP config: ( definir o padrão de cada objeto e nome do app)
1. Adicionar em config.json um novo objeto com o tema do app.
2. Adicionar em notifications o novo app.
3. Duplicar pasta em <root>/apps/base com o nome ${AppFlavor}. 
4. Adicionar Imagens em apps/
5. Adicionar Imagens em <root>/android/src/${APP_FLAVOR}/res/



