


CLI para criar uma nova "flavor" de Android
Esta CLI (Interface de Linha de Comando) foi desenvolvida em Go e possui o objetivo de facilitar a criação de uma nova "flavor" (sabor) para um aplicativo Android. Ela fornece um comando chamado create-android-flavor (ou new-flavor como um alias) que aceita uma série de flags para configurar a nova flavor.

Como usar
Para utilizar a CLI, execute o comando create-android-flavor seguido do caminho do diretório do projeto Android e das flags necessárias. A seguir estão as flags disponíveis para configuração:

--BUNDLE_ID: O ID do pacote do aplicativo Android (bundle ID).
--APP_FLAVOR: O nome da flavor a ser criada.
--BUILD_OUTPUT_TYPE: O tipo de saída da construção do aplicativo (por exemplo, "AAB").
--APP_KEY_ALIAS: O alias da chave do aplicativo.
--APP_KEY_PASSWORD: A senha da chave do aplicativo.
--APP_KEY_STORE_PASSWORD: A senha do arquivo de armazenamento da chave do aplicativo.
--APP_NAME: O nome do aplicativo.
--DEEP_LINKING_TAG: A tag de "deep linking" do aplicativo.
--PACKAGE_SRC: O pacote de origem do aplicativo.
--ICON_LAUNCHER_PATH: O caminho para o ícone de lançamento do aplicativo.
Funcionalidade
Ao executar o comando create-android-flavor, a CLI irá validar se o diretório do projeto Android fornecido é válido. Essa validação pode envolver a verificação de arquivos como o build.gradle ou outros arquivos relevantes para garantir que o diretório esteja correto e pronto para a criação da nova flavor.

Uma vez validado o diretório, a CLI irá prosseguir para criar a nova flavor. Essa criação pode envolver várias etapas, como:

Atualização do arquivo build.gradle para adicionar a nova flavor e suas configurações correspondentes.
Configuração do ID do pacote do aplicativo (bundle ID) para a nova flavor.
Configuração do tipo de saída da construção do aplicativo para a nova flavor.
Configuração das chaves de assinatura do aplicativo (alias da chave, senha da chave e senha do arquivo de armazenamento da chave) para a nova flavor.
Configuração do nome do aplicativo e da tag de "deep linking" para a nova flavor.
Configuração do pacote de origem e do ícone de lançamento do aplicativo para a nova flavor.
Após a conclusão dessas etapas, a nova flavor estará pronta para ser utilizada e compilada dentro do projeto Android.

Exemplo de Uso
Aqui está um exemplo de como utilizar a CLI para gerar uma nova flavor de aplicativo com um ícone e keystore personalizados:


```sh

create-android-flavor ./examples/android --BUNDLE_ID="com.example.facebook" --APP_FLAVOR="facebook" --BUILD_OUTPUT_TYPE="AAB" --APP_KEY_ALIAS="my-key-alias" --APP_KEY_PASSWORD="my-password" --APP_KEY_STORE_PASSWORD="my-app-keystore-password" --APP_NAME="facebook" --DEEP_LINKING_TAG="facebookApp-8574" --PACKAGE_SRC="com.example" --ICON_LAUNCHER_PATH="examples/

```