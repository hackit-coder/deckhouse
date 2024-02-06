# Администрирование платформы Deckhouse Kubernetes Platform

### Требования к ресурсам и окружению

{% alert %}
Выделяйте от 4 CPU / 8 ГБ RAM на инфраструктурные узлы. Для мастер-узлов и узлов мониторинга используйте быстрые диски.
{% endalert %}

Рекомендуются следующие минимальные ресурсы для инфраструктурных узлов в зависимости от их роли в кластере:
- **Мастер-узел** — 4 CPU, 8 ГБ RAM; быстрый диск с не менее чем 400 IOPS.  
- **Frontend-узел** — 2 CPU, 4 ГБ RAM;
- **Узел мониторинга** (для нагруженных кластеров) — 4 CPU, 8 ГБ RAM; быстрый диск.
- **Системный узел**:
  - 2 CPU, 4 ГБ RAM — если в кластере есть выделенные узлы мониторинга;
  - 4 CPU, 8 ГБ RAM, быстрый диск — если в кластере нет выделенных узлов мониторинга.

Примерный расчет ресурсов, необходимых для кластера:
- **Типовой кластер**: 3 мастер-узла, 2 frontend-узла, 2 системных узла. Такая конфигурация потребует **от 24 CPU и 48 ГБ RAM**, плюс быстрые диски с 400+ IOPS для мастер-узлов.
- **Кластер с повышенной нагрузкой** (с выделенными узлами мониторинга): 3 мастер-узла, 2 frontend-узла, 2 системных узла, 2 узла мониторинга. Такая конфигурация потребует **от 28 CPU и 64 ГБ RAM**, плюс быстрые диски с 400+ IOPS для мастер-узлов и узлов мониторинга.
- Для компонентов Deckhouse желательно выделить отдельный [storageClass](/documentation/v1/deckhouse-configure-global.html#parameters-storageclass) на быстрых дисках.
- Добавьте к этому ресурсы, необходимые для запуска полезной нагрузки.

## Доступ

Доступ к Deckhouse осуществляется через веб-интерфейс, который доступен по адресу https://deckhouse.io/. Для доступа к интерфейсу необходимо ввести логин и пароль, которые были получены при регистрации на сайте (уточнить).

Веб-интерфейс состоит из нескольких разделов, каждый из которых предназначен для выполнения определенных задач. Например, раздел “Dashboard” содержит информацию о состоянии кластера и его компонентах, раздел “Settings” позволяет настраивать параметры кластера и т. д.

Для доступа к определенным функциям Deckhouse может потребоваться наличие определенных прав доступа. Права доступа определяются на уровне ролей (roles), которые могут быть привязаны к конкретным пользователям или группам пользователей.

https://deckhouse.ru/documentation/v1/installing/configuration.html#

### Особенности конфигурации

**Мастер-узлы**

{% alert %}
В кластере должно быть три мастер-узла с быстрыми дисками 400+ IOPS.
{% endalert %}

Всегда используйте три мастер-узла — такое количество обеспечит отказоустойчивость и позволит безопасно выполнять обновление мастер-узлов. В большем числе мастер-узлов нет необходимости, а два узла не обеспечат кворума.

Конфигурация мастер-узлов для облачных кластеров настраивается в параметре [masterNodeGroup](/documentation/v1/modules/030-cloud-provider-aws/cluster_configuration.html#awsclusterconfiguration-masternodegroup).

Может быть полезно:
- [Как добавить мастер-узлы в облачном кластере...](/documentation/v1/modules/040-control-plane-manager/faq.html#как-добавить-master-узлы-в-облачном-кластере-single-master-в-multi-master)
- [Работа со статическими узлами...](/documentation/latest/modules/040-node-manager/#работа-со-статическими-узлами)

**Frontend-узлы**

{% alert %}
Выделите два или более frontend-узла.

Используйте inlet `LoadBalancer` для OpenStack и облачных сервисов, где возможен автоматический заказ балансировщика (Yandex Cloud, VK Cloud, Selectel Cloud, AWS, GCP, Azure и т. п.). Используйте inlet `HostPort` с внешним балансировщиком для bare metal или vSphere.
{% endalert %}

Frontend-узлы балансируют входящий трафик. На них работают Ingress-контроллеры. У [NodeGroup](/documentation/v1/modules/040-node-manager/cr.html#nodegroup) frontend-узлов установлен label `node-role.deckhouse.io/frontend`. Читайте подробнее про [выделение узлов под определенный вид нагрузки...](/documentation/v1/#выделение-узлов-под-определенный-вид-нагрузки)

Используйте более одного frontend-узла. Frontend-узлы должны выдерживать трафик при отказе как минимум одного frontend-узла.

Например, если в кластере два frontend-узла, то каждый frontend-узел должен справляться со всей нагрузкой на кластер в случае, если второй выйдет из строя. Если в кластере три frontend-узла, то каждый frontend-узел должен выдерживать увеличение нагрузки как минимум в полтора раза.

Выберите [тип inlet'а](/documentation/v1/modules/402-ingress-nginx/cr.html#ingressnginxcontroller-v1-spec-inlet) (он определяет способ поступления трафика).  

При развертывании кластера с помощью Deckhouse в облачной инфраструктуре, в которой поддерживается заказ балансировщиков (например, решения на базе OpenStack, сервисы Yandex Cloud, VK Cloud, Selectel Cloud, AWS, GCP, Azure и т. п.), используйте inlet `LoadBalancer` или `LoadBalancerWithProxyProtocol`.

В средах, в которых автоматический заказ балансировщиков недоступен (в bare-metal-кластерах, vSphere, некоторых решениях на базе OpenStack), используйте inlet `HostPort` или `HostPortWithProxyProtocol`. В этом случае можно либо добавить несколько A&#8209;записей в DNS для соответствующего домена, либо использовать внешний сервис балансировки нагрузки (например, взять решения от Cloudflare, Qrator или настроить metallb).

{% alert level="warning" %}
Inlet `HostWithFailover` подходит для кластеров с одним frontend-узлом. Он позволяет сократить время недоступности Ingress-контроллера при его обновлении. Такой тип inlet'а подойдет, например, для важных сред разработки, но **не рекомендуется для production**.
{% endalert %}

Алгоритм выбора inlet'а:

![Алгоритм выбора inlet'а]({{ assets["guides/going_to_production/ingress-inlet-ru.svg"].digest_path }})

**Узлы мониторинга**

{% alert %}
Для нагруженных кластеров выделите два узла мониторинга с быстрыми дисками.
{% endalert %}

Узлы мониторинга служат для запуска Grafana, Prometheus и других компонентов мониторинга. У [NodeGroup](/documentation/v1/modules/040-node-manager/cr.html#nodegroup) узлов мониторинга установлен label `node-role.deckhouse.io/monitoring`.

В нагруженных кластерах со множеством алертов и большими объемами метрик под мониторинг рекомендуется выделить отдельные узлы. Если этого не сделать, компоненты мониторинга будут размещены на [системных узлах](#системные-узлы).

При выделении узлов под мониторинг важно, чтобы на них были быстрые диски. Для этого можно привязать `storageClass` на быстрых дисках ко всем компонентам Deckhouse (глобальный параметр [storageClass](/documentation/v1/deckhouse-configure-global.html#parameters-storageclass)) или выделить отдельный `storageClass` только для компонентов мониторинга (параметры [storageClass](/documentation/v1/modules/300-prometheus/configuration.html#parameters-storageclass) и [longtermStorageClass](/documentation/v1/modules/300-prometheus/configuration.html#parameters-longtermstorageclass) модуля `prometheus`).

**Системные узлы**

{% alert %}
Выделите два системных узла.
{% endalert %}

Системные узлы предназначены для запуска модулей Deckhouse. У [NodeGroup](/documentation/v1/modules/040-node-manager/cr.html#nodegroup) системных узлов установлен label `node-role.deckhouse.io/system`.

Выделите два системных узла. В этом случае модули Deckhouse будут работать на них, не пересекаясь с пользовательскими приложениями кластера. Читайте подробнее про [выделение узлов под определенный вид нагрузки...](/documentation/v1/#выделение-узлов-под-определенный-вид-нагрузки).

Компонентам Deckhouse желательно выделить быстрые диски (глобальный параметр [storageClass](/documentation/v1/deckhouse-configure-global.html#parameters-storageclass)).

## Подключение

Чтобы подключить Deckhouse к своему кластеру Kubernetes, необходимо выполнить следующие шаги:

1. Скачайте и установите Helm. Helm - это инструмент для управления пакетами в Kubernetes. Он позволяет устанавливать, обновлять и удалять приложения в кластере.

1. Создайте новый Helm-релиз. Для этого необходимо создать файл с описанием релиза и значениями параметров. Файл должен иметь расширение .yaml и содержать информацию о том, какой компонент должен быть установлен, какие параметры у него должны быть и т.д.

2. Запустите Helm-релиз с помощью команды `helm install`. Эта команда принимает на вход файл с описанием Helm-релиза и устанавливает указанный компонент в кластере Kubernetes.
После установки компонента необходимо проверить, что он работает корректно. Для этого можно использовать стандартные инструменты Kubernetes, такие как kubectl, или специальные инструменты, предоставляемые Deckhouse.

(расписать подробно)

## Аутентификация

Аутентификация в Deckhouse Kubernetes Platform - это процесс подтверждения того, что пользователь является тем, за кого он себя выдает. Это необходимо для защиты данных и обеспечения безопасности системы.

В Deckhouse Kubernetes Platfor используется аутентификация на основе токенов. Это означает, что каждый пользователь получает уникальный токен, который он должен использовать для доступа к ресурсам системы. Токен содержит информацию о пользователе и сроке его действия.

Когда пользователь пытается получить доступ к ресурсу, он отправляет свой токен в заголовке запроса. Сервер проверяет токен и, если он действителен, разрешает доступ к ресурсу. Если токен недействителен, доступ отклоняется.

Важно следить за сроком действия токенов и обновлять их по мере необходимости. Также необходимо обеспечивать безопасность токенов, чтобы они не попали в руки злоумышленников.

https://deckhouse.ru/documentation/v1/modules/150-user-authn/

Для аутентификации в Deckhouse Kubernetes Platform используется модуль **User AuthN**. Этот модуль использует OAuth 2.0 для авторизации пользователей и предоставляет API для работы с аутентификационными токенами.

Модуль **User AuthN** состоит из следующих компонентов:

* Auth Server - сервер аутентификации, который обрабатывает запросы на аутентификацию и выдает аутентификационные токены.

* Resource Server - сервер ресурсов, который предоставляет доступ к ресурсам платформы только авторизованным пользователям.

* Client - клиент, который отправляет запросы на аутентификацию на сервер аутентификации и получает аутентификационные токены.

Процесс аутентификации выглядит следующим образом:

* Пользователь отправляет запрос на аутентификацию на сервер аутентификации.

(сюда сценарий)

* Сервер аутентификации проверяет учетные данные пользователя и, если они верны, генерирует аутентификационный токен.

(сюда сценарий)

* Сервер аутентификации возвращает токен пользователю.

(сюда сценарий)

* Пользователь передает токен на сервер ресурсов при запросе доступа к ресурсам.

(сюда сценарий)

* Сервер ресурсов проверяет токен и, если он действителен, предоставляет доступ к запрашиваемым ресурсам.

(сюда сценарий)

Таким образом, модуль User AuthN обеспечивает надежную аутентификацию пользователей и защиту их данных.

https://deckhouse.ru/documentation/v1/modules/150-user-authn/faq.html

https://deckhouse.ru/documentation/v1/modules/150-user-authn/usage.html

### Пример конфигурации модуля аутентификации

Представленный YAML-файл описывает конфигурацию модуля `user-authn` в системе Deckhouse и содержит настройки для генератора `kubeconfig` с публикацией API.

```yaml
apiVersion: deckhouse.io/v1alpha1
kind: ModuleConfig
metadata:
  name: user-authn
spec:
  version: 1
  enabled: true
  settings:
    kubeconfigGenerator:
    - id: direct
      masterURI: https://159.89.5.247:6443
      description: "Direct access to kubernetes API"
    publishAPI:
      enable: true
```

### Примеры настройки провайдера

#### GitHub

Представленный YAML-файл описывает конфигурацию провайдера Dex, который использует GitHub для аутентификации. Провайдер имеет имя "github" и тип "Github". В разделе "github" указаны клиентский идентификатор и секрет.

```yaml
apiVersion: deckhouse.io/v1
kind: DexProvider
metadata:
  name: github
spec:
  type: Github
  displayName: My Company GitHub
  github:
    clientID: plainstring
    clientSecret: plainstring
```

1. В организации GitHub создайте новое приложение.

Для этого перейдите в `Settings` -> `Developer settings` -> `OAuth Aps` -> `Register a new OAuth application` и в качестве `Authorization callback URL` указажите адрес `https://dex.<modules.publicDomainTemplate>/callback`.

2. Полученные `Client ID` и `Client Secret` укажите в custom resource [DexProvider](cr.html#dexprovider).

В том случае, если организация GitHub находится под управлением клиента, перейдите в `Settings` -> `Applications` -> `Authorized OAuth Apps` -> `<name of created OAuth App>` и запросите подтверждение нажатием на `Send Request`. 

3. Попросите клиента подтвердить запрос, который придет к нему на email.

#### GitLab

Представленный YAML-файл описывает конфигурацию провайдера Dex, который использует GitLab для аутентификации. Провайдер имеет имя "gitlab" и тип "Gitlab". В разделе "gitlab" указан базовый URL GitLab, клиентский идентификатор и секрет, а также группы, которые будут иметь доступ к аутентификации через этого провайдера.

```yaml
apiVersion: deckhouse.io/v1
kind: DexProvider
metadata:
  name: gitlab
spec:
  type: Gitlab
  displayName: Dedicated Gitlab
  gitlab:
    baseURL: https://gitlab.example.com
    clientID: plainstring
    clientSecret: plainstring
    groups:
    - administrators
    - users
```

1. В GitLab проекта создайте новое приложение.

Для этого сделайте следующие шаги:
* **self-hosted**: перейдите в `Admin area` -> `Application` -> `New application` и в качестве `Redirect URI (Callback url)` укажите адрес `https://dex.<modules.publicDomainTemplate>/callback`, scopes выберите: `read_user`, `openid`;
* **cloud gitlab.com**: под главной учетной записью проекта перейдите в `User Settings` -> `Application` -> `New application` и в качестве `Redirect URI (Callback url)` укажите адрес `https://dex.<modules.publicDomainTemplate>/callback`, scopes выберите: `read_user`, `openid`;
* (для GitLab версии 16 и выше) включите опцию `Trusted`/`Trusted applications are automatically authorized on Gitlab OAuth flow` при создании приложения.

2. Полученные `Application ID` и `Secret` укажите в custom resource [DexProvider](cr.html#dexprovider).

#### Atlassian Crowd

```yaml
apiVersion: deckhouse.io/v1
kind: DexProvider
metadata:
  name: crowd
spec:
  type: Crowd
  displayName: Crowd
  crowd:
    baseURL: https://crowd.example.com/crowd
    clientID: plainstring
    clientSecret: plainstring
    enableBasicAuth: true
    groups:
    - administrators
    - users
```

В соответствующем проекте Atlassian Crowd необходимо создать новое `Generic`-приложение.

Для этого необходимо перейти в `Applications` -> `Add application`.

Полученные `Application Name` и `Password` необходимо указать в custom resource [DexProvider](cr.html#dexprovider).

Группы CROWD указываются в lowercase-формате для custom resource `DexProvider`.

#### Bitbucket Cloud

```yaml
apiVersion: deckhouse.io/v1
kind: DexProvider
metadata:
  name: gitlab
spec:
  type: BitbucketCloud
  displayName: Bitbucket
  bitbucketCloud:
    clientID: plainstring
    clientSecret: plainstring
    includeTeamGroups: true
    teams:
    - administrators
    - users
```

Для настройки аутентификации необходимо в Bitbucket в меню команды создать нового OAuth consumer.

Для этого необходимо перейти в `Settings` -> `OAuth consumers` -> `New application` и в качестве `Callback URL` указать адрес `https://dex.<modules.publicDomainTemplate>/callback`, разрешить доступ для `Account: Read` и `Workspace membership: Read`.

Полученные `Key` и `Secret` необходимо указать в custom resource [DexProvider](cr.html#dexprovider).

#### OIDC (OpenID Connect)

```yaml
apiVersion: deckhouse.io/v1
kind: DexProvider
metadata:
  name: okta
spec:
  type: OIDC
  displayName: My Company Okta
  oidc:
    issuer: https://my-company.okta.com
    clientID: plainstring
    clientSecret: plainstring
    insecureSkipEmailVerified: true
    getUserInfo: true
```

Аутентификация через OIDC-провайдера требует регистрации клиента (или создания приложения). Сделайте это по документации вашего провайдера (например, [Okta](https://help.okta.com/en-us/Content/Topics/Apps/Apps_App_Integration_Wizard_OIDC.htm), [Keycloak](https://www.keycloak.org/docs/latest/server_admin/index.html#proc-creating-oidc-client_server_administration_guide), [Gluu](https://gluu.org/docs/gluu-server/4.4/admin-guide/openid-connect/#manual-client-registration)).

Полученные в ходе выполнения инструкции `clientID` и `clientSecret` необходимо указать в custom resource [DexProvider](cr.html#dexprovider).

#### LDAP

```yaml
apiVersion: deckhouse.io/v1
kind: DexProvider
metadata:
  name: active-directory
spec:
  type: LDAP
  displayName: Active Directory
  ldap:
    host: ad.example.com:636
    insecureSkipVerify: true

    bindDN: cn=Administrator,cn=users,dc=example,dc=com
    bindPW: admin0!

    usernamePrompt: Email Address

    userSearch:
      baseDN: cn=Users,dc=example,dc=com
      filter: "(objectClass=person)"
      username: userPrincipalName
      idAttr: DN
      emailAttr: userPrincipalName
      nameAttr: cn

    groupSearch:
      baseDN: cn=Users,dc=example,dc=com
      filter: "(objectClass=group)"
      userMatchers:
      - userAttr: DN
        groupAttr: member
      nameAttr: cn
```

Для настройки аутентификации необходимо завести в LDAP read-only-пользователя (service account).

Полученные путь до пользователя и пароль необходимо указать в полях `bindDN` и `bindPW` custom resource [DexProvider](cr.html#dexprovider).
1. Если в LDAP настроен анонимный доступ на чтение, настройки можно не указывать.
2. В поле `bindPW` необходимо указывать пароль в plain-виде. Стратегии с передачей хэшированных паролей не предусмотрены.

### Настройка OAuth2-клиента в Dex для подключения приложения

Данный вариант настройки подходит приложениям, которые имеют возможность использовать oauth2-аутентификацию самостоятельно, без помощи oauth2-proxy.
Чтобы позволить подобным приложениям взаимодействовать с Dex, используется custom resource [`DexClient`](cr.html#dexclient).

```yaml
apiVersion: deckhouse.io/v1
kind: DexClient
metadata:
  name: myname
  namespace: mynamespace
spec:
  redirectURIs:
  - https://app.example.com/callback
  - https://app.example.com/callback-reserve
  allowedGroups:
  - Everyone
  - admins
  trustedPeers:
  - opendistro-sibling
```

После создания такого ресурса в Dex будет зарегистрирован клиент с идентификатором (**clientID**) `dex-client-myname@mynamespace`.

Пароль для доступа к клиенту (**clientSecret**) будет сохранен в Secret'е:
{% raw %}

```yaml
apiVersion: v1
kind: Secret
metadata:
  name: dex-client-myname
  namespace: mynamespace
type: Opaque
data:
  clientSecret: c2VjcmV0
```

### Пример создания статического пользователя

Придумайте пароль и укажите его хэш-сумму в поле `password`.

Для вычисления хэш-суммы пароля воспользуйтесь командой:

```shell
echo "$password" | htpasswd -inBC 10 "" | tr -d ':\n' | sed 's/$2y/$2a/'
```

Также можно воспользоваться [онлайн-сервисом](https://bcrypt-generator.com/).

{% raw %}

```yaml
apiVersion: deckhouse.io/v1
kind: User
metadata:
  name: admin
spec:
  email: admin@yourcompany.com
  password: $2a$10$etblbZ9yfZaKgbvysf1qguW3WULdMnxwWFrkoKpRH1yeWa5etjjAa
  ttl: 24h
```

{% endraw %}

#### Пример добавления статического пользователя в группу

{% raw %}

```yaml
apiVersion: deckhouse.io/v1alpha1
kind: Group
metadata:
  name: admins
spec:
  name: admins
  members:
    - kind: User
      name: admin
```


https://deckhouse.ru/documentation/v1/modules/140-user-authz/usage.html


#### Пример `ClusterAuthorizationRule`

```yaml
apiVersion: deckhouse.io/v1
kind: ClusterAuthorizationRule
metadata:
  name: test-rule
spec:
  subjects:
  - kind: User
    name: some@example.com
  - kind: ServiceAccount
    name: gitlab-runner-deploy
    namespace: d8-service-accounts
  - kind: Group
    name: some-group-name
  accessLevel: PrivilegedUser
  portForwarding: true
  # Опция доступна только при включенном режиме enableMultiTenancy (версия Enterprise Edition).
  allowAccessToSystemNamespaces: false
  # Опция доступна только при включенном режиме enableMultiTenancy (версия Enterprise Edition).
  namespaceSelector:
    labelSelector:
      matchExpressions:
      - key: stage
        operator: In
        values:
        - test
        - review
      matchLabels:
        team: frontend
```

### Создание пользователя

В Kubernetes есть две категории пользователей:

* ServiceAccount'ы, учет которых ведет сам Kubernetes через API.
* Остальные пользователи, учет которых ведет не сам Kubernetes, а некоторый внешний софт, который настраивает администратор кластера, — существует множество механизмов аутентификации и, соответственно, множество способов заводить пользователей. В настоящий момент поддерживаются два способа аутентификации:
  * через модуль [user-authn](../../modules/150-user-authn/);
  * с помощью сертификатов.

При выпуске сертификата для аутентификации нужно указать в нем имя (`CN=<имя>`), необходимое количество групп (`O=<группа>`) и подписать его с помощью корневого CA-кластера. Именно этим механизмом вы аутентифицируетесь в кластере, когда, например, используете kubectl на bastion-узле.

#### Создание ServiceAccount для сервера и предоставление ему доступа

Создание ServiceAccount с доступом к Kubernetes API может потребоваться, например, при настройке развертывания приложений через CI-системы.  

1. Создайте ServiceAccount, например в namespace `d8-service-accounts`:

   ```shell
   kubectl create -f - <<EOF
   apiVersion: v1
   kind: ServiceAccount
   metadata:
     name: gitlab-runner-deploy
     namespace: d8-service-accounts
   ---
   apiVersion: v1
   kind: Secret
   metadata:
     name: gitlab-runner-deploy-token
     namespace: d8-service-accounts
     annotations:
       kubernetes.io/service-account.name: gitlab-runner-deploy
   type: kubernetes.io/service-account-token
   EOF
   ```

1. Дайте необходимые ServiceAccount права (используя custom resource [ClusterAuthorizationRule](cr.html#clusterauthorizationrule)):

   ```shell
   kubectl create -f - <<EOF
   apiVersion: deckhouse.io/v1
   kind: ClusterAuthorizationRule
   metadata:
     name: gitlab-runner-deploy
   spec:
     subjects:
     - kind: ServiceAccount
       name: gitlab-runner-deploy
       namespace: d8-service-accounts
     accessLevel: SuperAdmin
     # Опция доступна только при включенном режиме enableMultiTenancy (версия Enterprise Edition).
     allowAccessToSystemNamespaces: true      
   EOF
   ```

   Если в конфигурации Deckhouse включен режим мультитенантности (параметр [enableMultiTenancy](configuration.html#parameters-enablemultitenancy), доступен только в Enterprise Edition), настройте доступные для ServiceAccount пространства имен (параметр [namespaceSelector](cr.html#clusterauthorizationrule-v1-spec-namespaceselector)).

1. Определите значения переменных (они будут использоваться далее), выполнив следующие команды (**подставьте свои значения**):

   ```shell
   export CLUSTER_NAME=my-cluster
   export USER_NAME=gitlab-runner-deploy.my-cluster
   export CONTEXT_NAME=${CLUSTER_NAME}-${USER_NAME}
   export FILE_NAME=kube.config
   ```

1. Сгенерируйте секцию `cluster` в файле конфигурации kubectl:

   Используйте один из следующих вариантов доступа к API-серверу кластера:

   * Если есть прямой доступ до API-сервера:
     1. Получите сертификат CA кластера Kubernetes:

        ```shell
        kubectl get cm kube-root-ca.crt -o jsonpath='{ .data.ca\.crt }' > /tmp/ca.crt
        ```

     1. Сгенерируйте секцию `cluster` (используется IP-адрес API-сервера для доступа):

        ```shell
        kubectl config set-cluster $CLUSTER_NAME --embed-certs=true \
          --server=https://$(kubectl get ep kubernetes -o json | jq -rc '.subsets[0] | "\(.addresses[0].ip):\(.ports[0].port)"') \
          --certificate-authority=/tmp/ca.crt \
          --kubeconfig=$FILE_NAME
        ```

   * Если прямого доступа до API-сервера нет, то используйте один следующих вариантов:
      * включите доступ к API-серверу через Ingress-контроллер (параметр [publishAPI](../150-user-authn/configuration.html#parameters-publishapi)), и укажите адреса с которых будут идти запросы (параметр [whitelistSourceRanges](../150-user-authn/configuration.html#parameters-publishapi-whitelistsourceranges));
      * укажите адреса с которых будут идти запросы в отдельном Ingress-контроллере (параметр [acceptRequestsFrom](../402-ingress-nginx/cr.html#ingressnginxcontroller-v1-spec-acceptrequestsfrom)).

   * Если используется непубличный CA:

     1. Получите сертификат CA из Secret'а с сертификатом, который используется для домена `api.%s`:

        ```shell
        kubectl -n d8-user-authn get secrets -o json \
          $(kubectl -n d8-user-authn get ing kubernetes-api -o jsonpath="{.spec.tls[0].secretName}") \
          | jq -rc '.data."ca.crt" // .data."tls.crt"' \
          | base64 -d > /tmp/ca.crt
        ```

     2. Сгенерируйте секцию `cluster` (используется внешний домен и CA для доступа):

        ```shell
        kubectl config set-cluster $CLUSTER_NAME --embed-certs=true \
          --server=https://$(kubectl -n d8-user-authn get ing kubernetes-api -ojson | jq '.spec.rules[].host' -r) \
          --certificate-authority=/tmp/ca.crt \
          --kubeconfig=$FILE_NAME
        ```

   * Если используется публичный CA. Сгенерируйте секцию `cluster` (используется внешний домен для доступа):

     ```shell
     kubectl config set-cluster $CLUSTER_NAME \
       --server=https://$(kubectl -n d8-user-authn get ing kubernetes-api -ojson | jq '.spec.rules[].host' -r) \
       --kubeconfig=$FILE_NAME
     ```

1. Сгенерируйте секцию `user` с токеном из Secret'а ServiceAccount в файле конфигурации kubectl:

   ```shell
   kubectl config set-credentials $USER_NAME \
     --token=$(kubectl -n d8-service-accounts get secret gitlab-runner-deploy-token -o json |jq -r '.data["token"]' | base64 -d) \
     --kubeconfig=$FILE_NAME
   ```

1. Сгенерируйте контекст в файле конфигурации kubectl:

   ```shell
   kubectl config set-context $CONTEXT_NAME \
     --cluster=$CLUSTER_NAME --user=$USER_NAME \
     --kubeconfig=$FILE_NAME
   ```

1. Установите сгенерированный контекст как используемый по умолчанию в файле конфигурации kubectl:

   ```shell
   kubectl config use-context $CONTEXT_NAME --kubeconfig=$FILE_NAME
   ```

##### Создание пользователя с помощью клиентского сертификата

* Получите корневой сертификат кластера (ca.crt и ca.key).
* Сгенерируйте ключ пользователя:

  ```shell
  openssl genrsa -out myuser.key 2048
  ```

* Создайте CSR, где укажите, что требуется пользователь `myuser`, который состоит в группах `mygroup1` и `mygroup2`:

  ```shell
  openssl req -new -key myuser.key -out myuser.csr -subj "/CN=myuser/O=mygroup1/O=mygroup2"
  ```

* Подпишите CSR корневым сертификатом кластера:

  ```shell
  openssl x509 -req -in myuser.csr -CA ca.crt -CAkey ca.key -CAcreateserial -out myuser.crt -days 10
  ```

* Теперь полученный сертификат можно указывать в конфиг-файле:

  ```shell
  cat << EOF
  apiVersion: v1
  clusters:
  - cluster:
      certificate-authority-data: $(cat ca.crt | base64 -w0)
      server: https://<хост кластера>:6443
    name: kubernetes
  contexts:
  - context:
      cluster: kubernetes
      user: myuser
    name: myuser@kubernetes
  current-context: myuser@kubernetes
  kind: Config
  preferences: {}
  users:
  - name: myuser
    user:
      client-certificate-data: $(cat myuser.crt | base64 -w0)
      client-key-data: $(cat myuser.key | base64 -w0)
  EOF
  ```

**Предоставление доступа созданному пользователю**

Для предоставления доступа созданному пользователю создайте `ClusterAuthorizationRule`.

Пример `ClusterAuthorizationRule`:

```yaml
apiVersion: deckhouse.io/v1
kind: ClusterAuthorizationRule
metadata:
  name: myuser
spec:
  subjects:
  - kind: User
    name: myuser
  accessLevel: PrivilegedUser
  portForwarding: true
```

### Настройка `kube-apiserver` для работы в режиме multi-tenancy

Режим multi-tenancy, позволяющий ограничивать доступ к namespace, включается параметром [enableMultiTenancy](configuration.html#parameters-enablemultitenancy) модуля.

Работа в режиме multi-tenancy требует включения [плагина авторизации Webhook](https://kubernetes.io/docs/reference/access-authn-authz/webhook/) и выполнения настройки `kube-apiserver`. Все необходимые для работы режима multi-tenancy действия **выполняются автоматически** модулем [control-plane-manager](../../modules/040-control-plane-manager/), никаких ручных действий не требуется.

Изменения манифеста `kube-apiserver`, которые произойдут после включения режима multi-tenancy:

* исправление аргумента `--authorization-mode`. Перед методом RBAC добавится метод Webhook (например — `--authorization-mode=Node,Webhook,RBAC`);
* добавление аргумента `--authorization-webhook-config-file=/etc/kubernetes/authorization-webhook-config.yaml`;
* добавление `volumeMounts`:

  ```yaml
  - name: authorization-webhook-config
    mountPath: /etc/kubernetes/authorization-webhook-config.yaml
    readOnly: true
  ```

* добавление `volumes`:

  ```yaml
  - name: authorization-webhook-config
    hostPath:
      path: /etc/kubernetes/authorization-webhook-config.yaml
      type: FileOrCreate
  ```

### Как проверить, что у пользователя есть доступ?

Необходимо выполнить следующую команду, в которой будут указаны:

* `resourceAttributes` (как в RBAC) — к чему мы проверяем доступ;
* `user` — имя пользователя;
* `groups` — группы пользователя.

> При совместном использовании с модулем `user-authn` группы и имя пользователя можно посмотреть в логах Dex — `kubectl -n d8-user-authn logs -l app=dex` (видны только при авторизации).

```shell
cat  <<EOF | 2>&1 kubectl  create --raw  /apis/authorization.k8s.io/v1/subjectaccessreviews -f - | jq .status
{
  "apiVersion": "authorization.k8s.io/v1",
  "kind": "SubjectAccessReview",
  "spec": {
    "resourceAttributes": {
      "namespace": "",
      "verb": "watch",
      "version": "v1",
      "resource": "pods"
    },
    "user": "system:kube-controller-manager",
    "groups": [
      "Admins"
    ]
  }
}
EOF
```

В результате увидим, есть ли доступ и на основании какой роли:

```json
{
  "allowed": true,
  "reason": "RBAC: allowed by ClusterRoleBinding \"system:kube-controller-manager\" of ClusterRole \"system:kube-controller-manager\" to User \"system:kube-controller-manager\""
}
```

Если в кластере включен режим **multi-tenancy**, нужно выполнить еще одну проверку, чтобы убедиться, что у пользователя есть доступ в namespace:

```shell
cat  <<EOF | 2>&1 kubectl --kubeconfig /etc/kubernetes/deckhouse/extra-files/webhook-config.yaml create --raw / -f - | jq .status
{
  "apiVersion": "authorization.k8s.io/v1",
  "kind": "SubjectAccessReview",
  "spec": {
    "resourceAttributes": {
      "namespace": "",
      "verb": "watch",
      "version": "v1",
      "resource": "pods"
    },
    "user": "system:kube-controller-manager",
    "groups": [
      "Admins"
    ]
  }
}
EOF
```

```json
{
  "allowed": false
}
```

Сообщение `allowed: false` значит, что webhook не блокирует запрос. В случае блокировки запроса webhook'ом вы получите, например, следующее сообщение:

```json
{
  "allowed": false,
  "denied": true,
  "reason": "making cluster scoped requests for namespaced resources are not allowed"
}
```

#### Настройка прав высокоуровневых ролей

Если требуется добавить прав для определенной [высокоуровневой роли](./#ролевая-модель), достаточно создать ClusterRole с аннотацией `user-authz.deckhouse.io/access-level: <AccessLevel>`.

Пример:

```yaml
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  annotations:
    user-authz.deckhouse.io/access-level: Editor
  name: user-editor
rules:
- apiGroups:
  - kuma.io
  resources:
  - trafficroutes
  - trafficroutes/finalizers
  verbs:
  - get
  - list
  - watch
  - create
  - update
  - patch
  - delete
- apiGroups:
  - flagger.app
  resources:
  - canaries
  - canaries/status
  - metrictemplates
  - metrictemplates/status
  - alertproviders
  - alertproviders/status
  verbs:
  - get
  - list
  - watch
  - create
  - update
  - patch
  - delete
```

### Основные параметры аутентификации

https://deckhouse.ru/documentation/v1/search.html?query=аутентификация
DexClient
Позволяет приложениям, поддерживающим DC-аутентификацию, взаимодействовать с Dex. После появления в кластере объекта DexClient : в Dex будет зарегистрирован клиент с идентификатором ( clientID ) dex-client-<NAME>@<NAMESPACE> где …

UpmeterRemoteWrite: username
spec.config.basicAuth.username
Имя для аутентификации.

UpmeterRemoteWrite: password
spec.config.basicAuth.password
Пароль для аутентификации.

UpmeterRemoteWrite: bearerToken
spec.config.bearerToken
Токен для аутентификации.

PrometheusRemoteWrite: password
spec.basicAuth.password
Пароль для аутентификации.

PrometheusRemoteWrite: username
spec.basicAuth.username
Имя пользователя для аутентификации.

Модуль deckhouse: настройки: basic
update.notification.auth.basic
Basic-аутентификация на webhook.

ClusterLogDestination: awsRegion
spec.elasticsearch.auth.awsRegion
Регион AWS для аутентификации.

ClusterLogDestination: strategy
spec.loki.auth.strategy
Используемый тип аутентификации.

ClusterLogDestination: token
spec.loki.auth.token
Токен для Bearer-аутентификации.

CustomAlertmanager: authIdentity
spec.internal.receivers.emailConfigs.authIdentity
Идентификатор, используемый для аутентификации.

CustomAlertmanager: authUsername
spec.internal.receivers.emailConfigs.authUsername
Имя пользователя, используемое для аутентификации.

CustomAlertmanager: corpID
spec.internal.receivers.wechatConfigs.corpID
Идентификатор корпорации для аутентификации.

ClusterLogDestination: user
spec.elasticsearch.auth.user
Имя пользователя, используемое при Basic-аутентификации.

ClusterLogDestination: user
spec.loki.auth.user
Имя пользователя, используемое при Basic-аутентификации.

CustomAlertmanager: basicAuth
spec.internal.receivers.opsgenieConfigs.httpConfig.basicAuth
Параметры базовой аутентификации клиента.

CustomAlertmanager: basicAuth
spec.internal.receivers.pagerdutyConfigs.httpConfig.basicAuth
Параметры базовой аутентификации клиента.

CustomAlertmanager: basicAuth
spec.internal.receivers.pushoverConfigs.httpConfig.basicAuth
Параметры базовой аутентификации клиента.

CustomAlertmanager: basicAuth
spec.internal.receivers.slackConfigs.httpConfig.basicAuth
Параметры базовой аутентификации клиента.

CustomAlertmanager: basicAuth
spec.internal.receivers.victoropsConfigs.httpConfig.basicAuth
Параметры базовой аутентификации клиента.

CustomAlertmanager: basicAuth
spec.internal.receivers.webhookConfigs.httpConfig.basicAuth
Параметры базовой аутентификации клиента.

CustomAlertmanager: basicAuth
spec.internal.receivers.wechatConfigs.httpConfig.basicAuth
Параметры базовой аутентификации клиента.

Prometheus-мониторинг: настройки: authURL
auth.externalAuthentication.authURL
URL сервиса аутентификации. Если пользователь прошел аутентификацию, сервис должен возвращать код ответа HTTP

Модуль istio: настройки: authURL
auth.externalAuthentication.authURL
URL сервиса аутентификации. Если пользователь прошел аутентификацию, сервис должен возвращать код ответа HTTP

Модуль cilium-hubble: настройки: authURL
auth.externalAuthentication.authURL
URL сервиса аутентификации. Если пользователь прошел аутентификацию, сервис должен возвращать код ответа HTTP

Модуль documentation: настройки: authURL
auth.externalAuthentication.authURL
URL сервиса аутентификации. Если пользователь прошел аутентификацию, сервис должен возвращать код ответа HTTP

Модуль openvpn: настройки: authURL
auth.externalAuthentication.authURL
URL сервиса аутентификации. Если пользователь прошел аутентификацию, сервис должен возвращать код ответа HTTP

Модуль dashboard: настройки: authURL
auth.externalAuthentication.authURL
URL сервиса аутентификации. Если пользователь прошел аутентификацию, сервис должен возвращать код ответа HTTP

Модуль upmeter: настройки: authURL
auth.status.externalAuthentication.authURL
URL сервиса аутентификации. Если пользователь прошел аутентификацию, сервис должен возвращать код ответа HTTP

Модуль upmeter: настройки: authURL
auth.webui.externalAuthentication.authURL
URL сервиса аутентификации. Если пользователь прошел аутентификацию, сервис должен возвращать код ответа HTTP

Модуль upmeter: настройки: authURL
smokeMini.auth.externalAuthentication.authURL
URL сервиса аутентификации. Если пользователь прошел аутентификацию, сервис должен возвращать код ответа HTTP

Prometheus-мониторинг: настройки: authSignInURL
auth.externalAuthentication.authSignInURL
URL, куда будет перенаправлен пользователь для прохождения аутентификации (если сервис аутентификации вернул код ответа HTTP, отличный от ).

Модуль istio: настройки: authSignInURL
auth.externalAuthentication.authSignInURL
URL, куда будет перенаправлен пользователь для прохождения аутентификации (если сервис аутентификации вернул код ответа HTTP, отличный от ).

Модуль cilium-hubble: настройки: authSignInURL
auth.externalAuthentication.authSignInURL
URL, куда будет перенаправлен пользователь для прохождения аутентификации (если сервис аутентификации вернул код ответа HTTP, отличный от ).

Модуль documentation: настройки: authSignInURL
auth.externalAuthentication.authSignInURL
URL, куда будет перенаправлен пользователь для прохождения аутентификации (если сервис аутентификации вернул код ответа HTTP, отличный от ).

Модуль openvpn: настройки: authSignInURL
auth.externalAuthentication.authSignInURL
URL, куда будет перенаправлен пользователь для прохождения аутентификации (если сервис аутентификации вернул код ответа HTTP, отличный от ).

Модуль dashboard: настройки: authSignInURL
auth.externalAuthentication.authSignInURL
URL, куда будет перенаправлен пользователь для прохождения аутентификации (если сервис аутентификации вернул код ответа HTTP, отличный от ).

Модуль upmeter: настройки: authSignInURL
auth.status.externalAuthentication.authSignInURL
URL, куда будет перенаправлен пользователь для прохождения аутентификации (если сервис аутентификации вернул код ответа HTTP, отличный от ).

Модуль upmeter: настройки: authSignInURL
auth.webui.externalAuthentication.authSignInURL
URL, куда будет перенаправлен пользователь для прохождения аутентификации (если сервис аутентификации вернул код ответа HTTP, отличный от ).

Модуль upmeter: настройки: authSignInURL
smokeMini.auth.externalAuthentication.authSignInURL
URL, куда будет перенаправлен пользователь для прохождения аутентификации (если сервис аутентификации вернул код ответа HTTP, отличный от ).

Prometheus-мониторинг: настройки: auth
auth
Опции, связанные с аутентификацией или авторизацией в приложении.

Модуль istio: настройки: auth
auth
Опции, связанные с аутентификацией или авторизацией в приложении.

Модуль openvpn: настройки: auth
auth
Опции, связанные с аутентификацией или авторизацией в приложении.

Модуль dashboard: настройки: auth
auth
Опции, связанные с аутентификацией или авторизацией в приложении.

ClusterLogDestination: password
spec.elasticsearch.auth.password
Закодированный в Base64 пароль для Basic-аутентификации.

ClusterLogDestination: sasl
spec.kafka.sasl
Конфигурация аутентификации SASL для взаимодействия с Kafka.

ClusterLogDestination: password
spec.loki.auth.password
Закодированный в Base64 пароль для Basic-аутентификации.

ClusterLogDestination: strategy
spec.elasticsearch.auth.strategy
Тип аутентификации — Basic или AWS .

CustomAlertmanager: cert
spec.internal.receivers.telegramConfigs.httpConfig.tlsConfig.cert
Сертификат клиента для представления при выполнении аутентификации клиента.

Модуль upmeter: настройки: whitelistSourceRanges
auth.status.whitelistSourceRanges
Список адресов в формате CIDR, которым разрешено проходить аутентификацию. Если параметр не указан, аутентификацию разрешено проходить без ограничения по IP-адресу.

Модуль upmeter: настройки: whitelistSourceRanges
auth.webui.whitelistSourceRanges
Список адресов в формате CIDR, которым разрешено проходить аутентификацию. Если параметр не указан, аутентификацию разрешено проходить без ограничения по IP-адресу.

Модуль upmeter: настройки: whitelistSourceRanges
smokeMini.auth.whitelistSourceRanges
Список адресов в формате CIDR, которым разрешено проходить аутентификацию. Если параметр не указан, аутентификацию разрешено проходить без ограничения по IP-адресу.

DexAuthenticator: whitelistSourceRanges
spec.whitelistSourceRanges
Список адресов в формате CIDR, которым разрешено проходить аутентификацию. Если параметр не указан, аутентификацию разрешено проходить без ограничения по IP-адресу.

Модуль user-authn: настройки: description
kubeconfigGenerator.description
Текстовое описание, содержащее информацию о том, чем этот метод аутентификации отличается от других.

DexClient: trustedPeers
spec.trustedPeers
ID клиентов, которым позволена cross-аутентификация. Подробнее…

DexProvider: promptType
spec.oidc.promptType
Определяет — должен ли Issuer запрашивать подтверждение и давать подсказки при аутентификации. По умолчанию будет запрошено подтверждение при первой аутентификации. Допустимые значения могут изменяться в зависимости от Issuer.

Модуль cilium-hubble: настройки: auth
auth
Опции, связанные с аутентификацией и авторизацией доступа к веб-интерфейсу Hubble.

Модуль documentation: настройки: auth
auth
Опции, связанные с аутентификацией и авторизацией доступа к веб-интерфейсу документации.

Модуль metallb: настройки: password
bgpPeers.password
Пароль для аутентификациии для роутеров, требующих TCP MD5 авторизованных сессий.

DexClient: redirectURIs
spec.redirectURIs
Список адресов, на которые допустимо редиректить Dex’у после успешного прохождения аутентификации.

ClusterLogDestination: mechanism
spec.kafka.sasl.mechanism
Механизм аутентификации SASL. Поддерживаются только PLAIN и SCRAM-подобные механизмы.

CustomAlertmanager: type
spec.internal.receivers.telegramConfigs.httpConfig.authorization.type
Устанавливает тип аутентификации. По умолчанию — Bearer. Basic будет вызывать ошибку.

Модуль openvpn: настройки: whitelistSourceRanges
auth.whitelistSourceRanges
Массив адресов в формате CIDR, которым разрешено проходить аутентификацию для доступа в OpenVPN.

Модуль dashboard: настройки: whitelistSourceRanges
auth.whitelistSourceRanges
Массив адресов в формате CIDR, которым разрешено проходить аутентификацию для доступа в dashboard.

Prometheus-мониторинг: настройки: externalAuthentication
auth.externalAuthentication
Параметры для подключения внешней аутентификации (используется механизм NGINX Ingress external-auth , работающий на основе модуля NGINX auth_request . Внешняя аутентификация включается автоматически, если включен модуль user-authn .

Модуль istio: настройки: externalAuthentication
auth.externalAuthentication
Параметры для подключения внешней аутентификации (используется механизм NGINX Ingress external-auth , работающий на основе модуля Nginx auth_request . Внешняя аутентификация включается автоматически, если включен модуль user-authn .

Модуль cilium-hubble: настройки: externalAuthentication
auth.externalAuthentication
Параметры для подключения внешней аутентификации (используется механизм NGINX Ingress external-auth , работающий на основе модуля Nginx auth_request . Внешняя аутентификация включается автоматически, если включен модуль user-authn .

Модуль documentation: настройки: externalAuthentication
auth.externalAuthentication
Параметры для подключения внешней аутентификации (используется механизм NGINX Ingress external-auth , работающий на основе модуля Nginx auth_request . Внешняя аутентификация включается автоматически, если включен модуль user-authn .

Модуль openvpn: настройки: externalAuthentication
auth.externalAuthentication
Параметры для подключения внешней аутентификации (используется механизм NGINX Ingress external-auth , работающий на основе модуля Nginx auth_request . Внешняя аутентификация включается автоматически, если включен модуль user-authn .

Модуль dashboard: настройки: externalAuthentication
auth.externalAuthentication
Параметры для подключения внешней аутентификации (используется механизм NGINX Ingress external-auth , работающий на основе модуля Nginx auth_request . Внешняя аутентификация включается автоматически, если включен модуль user-authn .

Модуль upmeter: настройки: externalAuthentication
auth.status.externalAuthentication
Параметры для подключения внешней аутентификации (используется механизм NGINX Ingress external-auth , работающий на основе модуля NGINX auth_request . Внешняя аутентификация включается автоматически, если включен модуль user-authn .

Модуль upmeter: настройки: externalAuthentication
auth.webui.externalAuthentication
Параметры для подключения внешней аутентификации (используется механизм NGINX Ingress external-auth , работающий на основе модуля NGINX auth_request . Внешняя аутентификация включается автоматически, если включен модуль user-authn .

CustomAlertmanager: basicAuth
spec.internal.receivers.telegramConfigs.httpConfig.basicAuth
Параметры базовой аутентификации клиента. Это взаимоисключающая опция с разделом Authorization. Если оба определены, BasicAuth имеет приоритет.

CustomAlertmanager: password
spec.internal.receivers.telegramConfigs.httpConfig.basicAuth.password
Secret, который содержит пароль для аутентификации. Secret должен находиться в пространстве имен d8-monitoring .

CustomAlertmanager: username
spec.internal.receivers.telegramConfigs.httpConfig.basicAuth.username
Secret, который содержит имя пользователя для аутентификации. Secret должен находиться в пространстве имен d8-monitoring .

DexProvider: displayName
spec.displayName
Имя провайдера, которое будет отображено на странице выбора провайдера для аутентификации. Если настроен всего один провайдер, страница выбора провайдера показываться не будет.

CustomAlertmanager: authPassword
spec.internal.receivers.emailConfigs.authPassword
Ключ Secret’а, содержащий пароль для аутентификации. Secret должен находиться в пространстве имен d8-monitoring .

Управление control plane: настройки: authn
apiserver.authn
Опциональные параметры аутентификации клиентов Kubernetes API. По умолчанию используются данные из ConfigMap, устанавливаемого модулем user-authn .

CustomAlertmanager: authSecret
spec.internal.receivers.emailConfigs.authSecret
Ключ Secret’а, содержащий CRAM-MD5-секрет для аутентификации. Secret должен находиться в пространстве имен d8-monitoring .

CustomAlertmanager: bearerTokenSecret
spec.internal.receivers.telegramConfigs.httpConfig.bearerTokenSecret
Secret, содержащий токен носителя, который будет использоваться клиентом для аутентификации. Secret должен находиться в пространстве имен d8-monitoring .

Prometheus-мониторинг: настройки: satisfyAny
auth.satisfyAny
Разрешает пройти только одну из аутентификаций. В комбинации с опцией whitelistSourceRanges позволяет считать авторизованными всех пользователей из указанных сетей без ввода логина и пароля.

Модуль istio: настройки: satisfyAny
auth.satisfyAny
Разрешает пройти только одну из аутентификаций. В комбинации с опцией whitelistSourceRanges позволяет считать авторизованными всех пользователей из указанных сетей без ввода логина и пароля.

CustomAlertmanager: bearerTokenSecret
spec.internal.receivers.opsgenieConfigs.httpConfig.bearerTokenSecret
Ключ Secret’а, содержащий bearer-токен, который будет использоваться клиентом для аутентификации. Secret должен находиться в пространстве имен d8-monitoring .

CustomAlertmanager: bearerTokenSecret
spec.internal.receivers.pagerdutyConfigs.httpConfig.bearerTokenSecret
Ключ Secret’а, содержащий bearer-токен, который будет использоваться клиентом для аутентификации. Secret должен находиться в пространстве имен d8-monitoring .

CustomAlertmanager: bearerTokenSecret
spec.internal.receivers.pushoverConfigs.httpConfig.bearerTokenSecret
Ключ Secret’а, содержащий bearer-токен, который будет использоваться клиентом для аутентификации. Secret должен находиться в пространстве имен d8-monitoring .

CustomAlertmanager: bearerTokenSecret
spec.internal.receivers.slackConfigs.httpConfig.bearerTokenSecret
Ключ Secret’а, содержащий bearer-токен, который будет использоваться клиентом для аутентификации. Secret должен находиться в пространстве имен d8-monitoring .

CustomAlertmanager: bearerTokenSecret
spec.internal.receivers.victoropsConfigs.httpConfig.bearerTokenSecret
Ключ Secret’а, содержащий bearer-токен, который будет использоваться клиентом для аутентификации. Secret должен находиться в пространстве имен d8-monitoring .

CustomAlertmanager: bearerTokenSecret
spec.internal.receivers.webhookConfigs.httpConfig.bearerTokenSecret
Ключ Secret’а, содержащий bearer-токен, который будет использоваться клиентом для аутентификации. Secret должен находиться в пространстве имен d8-monitoring .

CustomAlertmanager: bearerTokenSecret
spec.internal.receivers.wechatConfigs.httpConfig.bearerTokenSecret
Ключ Secret’а, содержащий bearer-токен, который будет использоваться клиентом для аутентификации. Secret должен находиться в пространстве имен d8-monitoring .

Модуль upmeter: настройки: externalAuthentication
smokeMini.auth.externalAuthentication
Параметры для подключения внешней аутентификации. Используется механизм NGINX Ingress external-auth , работающей на основе модуля NGINX auth_request .

Модуль dashboard: настройки: accessLevel
accessLevel
Уровень доступа в dashboard, если отключен модуль user-authn и не включена внешняя аутентификация ( externalAuthentication ). Возможные значения описаны в user-authz . По умолчанию используется уровень User . В случае использования модуля user-authn …

DexAuthenticator: allowedGroups
spec.allowedGroups
Группы, пользователям которых разрешено проходить аутентификацию. Дополнительно параметр помогает ограничить список групп до тех, которые несут для приложения полезную информацию. Например, в случае если у пользователя более групп, но приложению …

Модуль dashboard: настройки: useBearerTokens
auth.externalAuthentication.useBearerTokens
Токены авторизации. dashboard должен работать с Kubernetes API от имени пользователя (сервис аутентификации при этом должен обязательно возвращать в своих ответах HTTP-заголовок Authorization, в котором должен быть bearer-token — именно под этим …

### Авторизация

Авторизация в Deckhouse Kubernetes Platform позволяет пользователям получить доступ к своим данным и функциям, предоставляемым платформой.

Чтобы авторизоваться в Deckhouse Kubernetes Platform выполните следующие шаги:

1. Откройте веб-браузер на своем устройстве и перейдите на [официальный сайт](https://deckhouse.io/).

(Далее пример входа)

1. В правом верхнем углу экрана нажмите на кнопку **Войти**.

1. На открывшейся странице авторизации введите свой адрес электронной почты, который использовали при регистрации в Deckhouse Kubernetes Platform.

1. Введите пароль, соответствующий вашему адресу электронной почты. Убедитесь, что вводите пароль правильно и он соответствует требованиям безопасности.

1. Нажмите на кнопку **Авторизоваться**.
Если вы используете двухфакторную аутентификацию, вам потребуется ввести одноразовый пароль, который будет отправлен на ваш зарегистрированный адрес электронной почты или сгенерирован вашим аутентификационным приложением (например, Google Authenticator или Authy).
После успешной авторизации вы будете перенаправлены на главную страницу Deckhouse Kubernetes Platform, где сможете начать работу с платформой и своими данными.

Авторизацию в Deckhouse обеспечивает модуль авторизации (Authentication Module, Auth). Auth модуль обеспечивает механизмы идентификации, аутентификации и авторизации пользователей, а также управление доступом к ресурсам.

Deckhouse поддерживает несколько модулей авторизации, включая Basic Auth, OAuth2, OpenID Connect и другие. Выбор конкретного модуля зависит от требований приложения и предпочтений администратора.

**Basic Auth**
Basic Authentication (Обычная Аутентификация) - это простой и широко распространенный метод аутентификации в HTTP. Он использует отправку пользователю имени пользователя и пароля в виде Base64-кодированных данных в заголовке HTTP-запроса.

(Сценарий)

**OAuth2**
OAuth 2.0 (Версия 2) - это открытый протокол авторизации, который позволяет одной части программного обеспечения (назовем его “клиент”) получить ограниченный доступ к защищенным ресурсам другой части (назовём её “сервер”) без необходимости передавать имя пользователя и пароль.

(Сценарий)

**OpenID Connect**
OpenID Connect (OIDC) - это расширение протокола OpenID, которое позволяет пользователям безопасно аутентифицироваться и получать информацию о своей учетной записи у поставщика удостоверений (IdP). OIDC обеспечивает аутентификацию и авторизацию пользователей, а также обмен информацией о пользователе между клиентом и поставщиком удостоверений.

(Сценарий)

Для настройки авторизации в Deckhouse, выполните следующие шаги:

1. Определите требования вашего приложения к авторизации и выберите подходящий модуль аутентификации (Basic Auth, OAuth2, OpenID Connect или другой).
1. Установите и настройте выбранный модуль аутентификации в соответствии с рекомендациями Deckhouse и вашей среды.
1. Настройте роли и политики доступа, чтобы определить, какие операции пользователи могут выполнять с ресурсами в вашем приложении.
1. Протестируйте систему авторизации, чтобы убедиться в ее корректной работе и соответствии требованиям безопасности.
1. При необходимости, настройте дополнительные параметры авторизации, такие как срок действия токенов, разрешения и т. д.
1. Обеспечьте обучение и поддержку пользователей по использованию системы авторизации в вашем приложении.

## Установка и сопровождение

Платформа Deckhouse Kubernetes Platform предоставляет инструменты для эффективного управления Kubernetes-кластерами и может быть установлена различными способами. В документации по установке Deckhouse (https://deckhouse.io/documentation/v1/installing/), можно найти подробное описание всех доступных вариантов установки, включая использование Helm-чартов, OCI-образов, а также инструкций для самостоятельной сборки из исходников. Помимо этого, для пользователей доступны подробные пошаговые инструкции по настройке и оптимизации работы платформы на различных инфраструктурах.

### Установка платформы

https://deckhouse.ru/documentation/v1/installing/

### Интеграция платформы с инфраструктурой

### Настройка платформы

## Обновление платформы

### Канал и режим обновлений

{% alert %}
Используйте канал обновлений `Early Access` или `Stable`. Установите [окно автоматических обновлений](/documentation/v1/modules/002-deckhouse/usage.html#конфигурация-окон-обновлений) или [ручной режим](/documentation/v1/modules/002-deckhouse/usage.html#ручное-подтверждение-обновлений).
{% endalert %}

Выберите [канал обновлений]( /documentation/v1/deckhouse-release-channels.html) и [режим обновлений](/documentation/v1/modules/002-deckhouse/configuration.html#parameters-releasechannel), который соответствует вашим ожиданиям. Чем стабильнее канал обновлений, тем позже до него доходит новая функциональность.

По возможности используйте разные каналы обновлений для кластеров. Для кластера разработки используйте менее стабильный канал обновлений, нежели для тестового или stage-кластера (pre-production-кластер).

Мы рекомендуем использовать канал обновлений `Early Access` или `Stable` для production-кластеров. Если в production-окружении более одного кластера, предпочтительно использовать для них разные каналы обновлений. Например, `Early Access` для одного, а `Stable` — для другого. Если использовать разные каналы обновлений по каким-либо причинам невозможно, рекомендуется устанавливать разные окна обновлений.

{% alert level="warning" %}
Даже в очень нагруженных и критичных кластерах не стоит отключать использование канала обновлений. Лучшая стратегия — плановое обновление. В инсталляциях Deckhouse, которые не обновлялись полгода или более, могут присутствовать ошибки. Как правило, эти ошибки давно устранены в новых версиях. В этом случае оперативно решить возникшую проблему будет непросто.
{% endalert %}

Управление [окнами обновлений](/documentation/v1/modules/002-deckhouse/configuration.html#parameters-update-windows) позволяет планово обновлять релизы Deckhouse в автоматическом режиме в периоды «затишья», когда нагрузка на кластер далека от пиковой.

#### Окна и каналы обновлений



### Высокая надежность и доступность
#### Going to production

https://deckhouse.ru/guides/production.html

### Управление ресурсами платформы
### Масштабирование платформы
## Безопасность платформы

https://deckhouse.ru/documentation/v1/deckhouse-overview.html#:~:text=%D0%B8%20%D1%83%D0%BF%D1%80%D0%B0%D0%B2%D0%BB%D0%B5%D0%BD%D0%B8%D0%B5%20%D1%80%D0%B5%D1%81%D1%83%D1%80%D1%81%D0%B0%D0%BC%D0%B8-,%D0%91%D0%B5%D0%B7%D0%BE%D0%BF%D0%B0%D1%81%D0%BD%D0%BE%D1%81%D1%82%D1%8C,-User%20authentication

### Дока по портам
### Дока по папкам для антивируса
### Мониторинг платформы

https://deckhouse.ru/documentation/v1/installing/configuration.html#

#### Уведомление о событиях мониторинга

{% alert %}
Настройте отправку алертов через [внутренний](/documentation/v1/modules/300-prometheus/faq.html#как-добавить-alertmanager) Alertmanager или подключите [внешний](/documentation/v1/modules/300-prometheus/faq.html#как-добавить-внешний-дополнительный-alertmanager).
{% endalert %}

Мониторинг будет работать сразу после установки Deckhouse, однако для production этого недостаточно. Чтобы получать уведомления об инцидентах, настройте [встроенный](/documentation/v1/modules/300-prometheus/faq.html#как-добавить-alertmanager) в Deckhouse Alertmanager или [подключите свой](/documentation/v1/modules/300-prometheus/faq.html#как-добавить-внешний-дополнительный-alertmanager) Alertmanager.

С помощью custom resource [CustomAlertmanager](/documentation/v1/modules/300-prometheus/cr.html#customalertmanager) можно настроить отправку уведомлений на [электронную почту](/documentation/v1/modules/300-prometheus/cr.html#customalertmanager-v1alpha1-spec-internal-receivers-emailconfigs), в [Slack](/documentation/v1/modules/300-prometheus/cr.html#customalertmanager-v1alpha1-spec-internal-receivers-slackconfigs), в [Telegram](/documentation/v1/modules/300-prometheus/usage.html#отправка-алертов-в-telegram), через [webhook](/documentation/v1/modules/300-prometheus/cr.html#customalertmanager-v1alpha1-spec-internal-receivers-webhookconfigs), а также другими способами.

#### Сбор логов

{% alert %}
[Настройте](/documentation/v1/modules/460-log-shipper/) централизованный сбор логов.
{% endalert %}

Настройте централизованный сбор логов с системных и пользовательских приложений с помощью модуля [log-shipper](/documentation/v1/modules/460-log-shipper/).

Достаточно создать custom resource с описанием того, *что нужно собирать*: [ClusterLoggingConfig](/documentation/v1/modules/460-log-shipper/cr.html#clusterloggingconfig) или [PodLoggingConfig](/documentation/v1/modules/460-log-shipper/cr.html#podloggingconfig); кроме того, необходимо создать custom resource с данными о том, *куда отправлять* собранные логи: [ClusterLogDestination](/documentation/v1/modules/460-log-shipper/cr.html#clusterlogdestination).

Дополнительная информация:
- [Пример для Grafana Loki](/documentation/v1/modules/460-log-shipper/examples.html#чтение-логов-из-всех-подов-кластера-и-направление-их-в-loki)
- [Пример для Logstash](/documentation/v1/modules/460-log-shipper/examples.html#простой-пример-logstash)
- [Пример для Splunk](/documentation/v1/modules/460-log-shipper/examples.html#пример-интеграции-со-splunk)

## Резервное копирование

{% alert %}
Настройте [резервное копирование etcd](/documentation/v1/modules/040-control-plane-manager/faq.html#как-сделать-бэкап-etcd). Напишите план восстановления.
{% endalert %}

Обязательно настройте [резервное копирование данных etcd](/documentation/v1/modules/040-control-plane-manager/faq.html#как-сделать-бэкап-etcd). Это будет ваш последний шанс на восстановление кластера в случае самых неожиданных событий. Храните резервные копии как можно *дальше* от кластера.  

Резервные копии не помогут, если они не работают или вы не знаете, как их использовать для восстановления. Рекомендуем составить [план восстановления на случай аварии](https://habr.com/ru/search/?q=%5BDRP%5D&target_type=posts&order=date) (Disaster Recovery Plan), содержащий конкретные шаги и команды по развертыванию кластера из резервной копии.

Этот план должен периодически актуализироваться и проверяться учебными тревогами.


### Журналирование платформы
## Контроль и управление доступом
### Разделение ресурсов платформы между пользователями
### Сообщество

{% alert %}
Следите за новостями проекта в [Telegram](https://t.me/deckhouse_ru).
{% endalert %}

Вступите в [сообщество](https://deckhouse.ru/community/about.html), чтобы быть в курсе важных изменений и новостей. Вы сможете общаться с людьми, занятыми общим делом. Это позволит избежать многих типичных проблем.

Команда Deckhouse знает, каких усилий требует организация работы production-кластера в Kubernetes. Мы будем рады, если Deckhouse позволит вам реализовать задуманное. Поделитесь своим опытом и вдохновите других на переход в Kubernetes.
