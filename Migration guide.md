# Инструкция по миграции c OpenShift на K8s

## Термины и определения
| Термины|  Описание|
| :--- | :--- |
| Виртуальная машина (машина) | Компьютерная система, эмулирующая возможности каких-либо вычислительных комплексов гостевых платформ (guest) на аппаратно-программном обеспечении хост-платформы (host) |
| Kubernetes (K8s) | Открытое программное обеспечение для автоматизации развертывания, масштабирования и управления контейнерными приложениями |
| K8s | Вариант поставки K8s, который включается в себя его основные компоненты, утилиты для развёртывания, CRI и набор операторов |
| Кластер | Состоит из набора машин (узлов - nodes), которые запускают контейнерные приложения. Кластер имеет, как минимум, один рабочий узел (worker node). В рабочих узлах размещены поды (pods), являющиеся компонентами приложения. Плоскость управления управляет рабочими узлами и подами в кластере. В промышленных средах плоскость управления обычно запускается на нескольких компьютерах, а кластер, как правило, развертывается на нескольких узлах, гарантируя отказоустойчивость и высокую надежность |
| Контейнер | Набор процессов, изолированный от остальной операционной системы и запускаемый с отдельного образа, который содержит все файлы, необходимые для их работы |
| CRI (container runtime interface) или среда выполнения контейнера | Программное обеспечение, которое создает контейнеры на основе образов контейнеров. Среда выступает в качестве посредника между контейнерами и операционной системой, предоставляя ресурсы, необходимые приложению, и управляя ими |
| CRI-O | Вариант исполнения (container runtime interface) помимо<b>docker</b>и containerd. Используется в K8s в качествеобработчика контейнеров. Позволяет загружать образы контейнеров из репозиториев, управлять ими и запускать их |
| Kubectl | Инструмент работы с командной строкой для управления объектами K8s |
| Kubelet | Агент, который работает на каждом узле в кластере K8s и следит за наличием подов на узлекластера K8s. Является связующим звеном между плоскостью управления (Control Plane) K8s и CRI на ноде кластера K8s. Представляет собой rpm-пакет, который устанавливается на ноды кластера K8s, в отличие от контейнеров, за развертывание которых отвечает CRI |
| Kubeadm | Инструмент командной строки, который устанавливает и настраивает различные компоненты кластера |
| Load Balancer (балансировщик нагрузки) | Программный или аппаратный продукт, который используется для распределения нагрузки по ресурсам кластера в зависимости от выбранной стратегии балансировки |
| Node (узел) | Виртуальная или физическая машина, на которой происходит запуск pods |
| Master node (главный узел) | Узел, на котором выполняется четыре основных процесса в кластере —<b>kube-apiserver</b>,<b>kube-controllermanager</b>,<b>kube-scheduler</b>и<b>etcd</b>. Для обеспечения отказоустойчивости используется нечетное количество узловот трех.На главных узлах не запускается пользовательская нагрузка |
| Worker node (рабочий узел) | Узел, на котором выполняются прикладные приложения |
| Pod | Группа из двух и более (<b>pause</b>и прикладная нагрузка) контейнеров с общим хранилищем, сетью и настройками для управления контейнерами. Pod является базовой единицей выполнения приложения K8s, самая маленькая и простая единица в объектной модели K8s, которую можно создать или развернуть |
| Persistent Volume | Абстракция K8s, позволяющая выделять дисковое/блочное пространство, существующее отдельно от pod и при необходимости монтировать это дисковое/блочное пространство в pod |
| Persistent Volume Claim | Используется для указания объектов<b>PersistentVolumes</b>в спецификации pod, так как объекты<b>PersistentVolumes</b>не могут быть указаны напрямую.Объекты<b>PersistentVolumeClaim</b>запрашивают определенный размер, режим доступа и класс хранилищ для объекта<b>PersistentVolume</b>. Если объект<b>PersistentVolume</b>, удовлетворяющий запросу, существует или может быть подготовлен, объект<b>PersistentVolumeClaim</b>связывается с необходимым объектом<b>PersistentVolume</b>. Кластер<b>Managed Service for Kubernetes</b>монтирует объект<b>PersistentVolumeClaim</b>в качестве тома для пода |
| Workloads (рабочая нагрузка) | Абстрактная модель группы pods в K8s. K8s подразделяет<b>workloads</b>на<b>Deployments</b>,<b>StatefulSets</b>,<b>DaemonSets</b>, задачи (jobs) и плановые задачи (cron jobs) |
| Kubernetes Namespace (Пространство имен) | Абстракция, позволяющая разделять вычислительные ресурсы и  разграничивать права доступа внутри одного кластера K8s между разными сотрудниками и/или командами |
| Label (метка) | Используются для группировки ресурсов в K8s. С помощью меток и селекторов можно управлять распределением pods по нодам кластера, а также предоставлять метаданные для приложений |
| Annotations (аннотации) | Позволяют ассоциировать с ресурсами произвольные метаданные, менять свойства или поведение объектов K8s, а также настраивать их взаимодействие с другими сервисами K8s |
| Controller (контроллер) | Процесс в K8s, который следит за состоянием ресурсов кластера K8s и, в случае, если текущее состояние отличается от заданного (желаемого), пытается привести состояние кластера к желаемому. |
| Replication controller (контроллер репликации) | Базовый тип контроллера. Определяет, как модули pod реплицируются по узлам кластера |
| Deployment (K8s) | Если экземпляр приложения не требует хранения состояния внутри себя и не важно какой именно экземпляр будет обрабатывать пришедший запрос, то для управления этими экземплярами отлично подойдет ресурс Deployment. Он позволяет запускать нескольких реплик приложения, а также управлять процессом обновления. Для управления количеством копий пода используется дополнительный ресурс<b>ReplicaSet</b> |
| DeploymentConfig | Oбъект, чуть усовершенствованный по сравнению с Deployment. Он использует триггеры и дополнительные объекты, такие как<b>BuildConfig</b> |
| ImageStream | Объект, который хранит в себе sha-хеш образов во внутреннем<b>registry OpenShift</b>и позволяет задать произвольный тег, независимо от того, что определено в registry. Стоит понимать, что это ещё один уровень абстракций K8s, а самих образов в<b>ImageStream</b>нет. Теги ссылаются либо на другой<b>ImageStream</b>, либо на Dockerimage. Помимо этого, с его помощью можно закешировать образ во внутреннем<b>registry</b>и не копировать его с внешнего контура |
| BuildConfig | Объект представляет собой сценарий сборки образа. Может содержать множество источников исходного кода (git, dockerfile и другие), стратегию сборки (Docker, S2I и другие). И назначение сборки —<b>Dockerimage</b>или<b>ImageStreamTag</b> |
| Service (сервис) | Абстракция, представляет собой логическое имя и IP-адрес, которые используются для обеспечения доступа к группе pods |
| ConfigMap (Конфигурация данных для приложения) | Объект K8s,который позволяет хранить конфигурационные данные в виде любого текста, например, конфигурационных файлов для nginx или пар ключ-значение. Он обычно используется для хранения настроек приложения, таких как параметры конфигурации, секреты или другие данные, необходимые для работы приложения. ConfigMap можно использовать для динамического изменения настроек приложения без необходимости перезапуска контейнера.Pods могут использовать<b>ConfigMap</b>в качестве переменных окружения или как файлы конфигурации в Volume. Переменные окружения позволяют использовать один и тот же Docker-образ для развертывания в различном окружении |
| Secret (секрет) | Аналогичен<b>ConfigMap</b>,но данные хранит в base64. Предназначен для хранения конфиденциальных данных — паролей, токенов, ключей, сертификатов и т.д.Secrets по умолчанию не шифруются в<b>etcd</b>, но есть возможность настроить шифрование при установке кластера.Использование secrets позволяет не включать секретные данные в код приложения. Secrets могут быть созданы независимо от pods, которые их используют, — это снижает риск раскрытия данных |
| Rpm-файл | Используется в RPM-like операционных системах. Файл RPM используется<b>Linux Package Manager</b>(пакетным менеджером) для установки, удаления, проверки, опроса и обновления программных пакетов) |
| Yum | Интерактивная автоматизированная программа обновления, которая может использоваться для сопровождения систем, использующих rpm. Альтернативой является DNF (Dandified Yum) - в отличие от yum, этот пакетный менеджер отличается увеличением скорости работы, низким потреблением памяти, предоставлением API для плагинов и интеграцией с другими приложениями |
| JWT (JSON Web Tokens) | Безопасный способ передачи информации между клиентом и сервером. JWT состоит из трех отдельных частей, разделенных точкой:<br>1. Заголовок: содержит метаданные о токене и используемом криптографическом алгоритме (HMAC SHA256, RSA и т.п.)<br> 2. Полезная нагрузка: фактические данные, которые содержит токен. Полезная нагрузка также известна как "утверждения" и может включать информацию о пользователе и дополнительных метаданных.<br>3. Подпись: криптографически защищенное доказательство (proof), которое подтверждает отправителя и гарантирует, что сообщение не было изменено во время передачи |
| SHA256 (Secure Hash Algorithm 256) | Один из наиболее распространенных алгоритмов хеширования, который принимает на вход сообщение любой длины и выдает на выходе 256-битный хеш-код |
| CoAP (Constrained Application Protocol) | Устройство, соединяющее сеть Интернет на основе протокола HTTР и сеть с ограниченными ресурсами, поддерживающую протокол CoAP. Прокси преобразует сообщения одной сети (протокол HTTP) в сообщения, понятные для другой (протокол CoAP) |
| DaemonSet | АбстракцияK8s , которая гарантирует, что на всех (или некоторых) узлах кластера запускается копия pod. По мере добавления узлов в кластер к ним добавляются pods. Когда узлы удаляются из кластера, эти pods удаляются. Удаление DaemonSet приводит к очистке созданных им модулей. Некоторые типичные варианты использования<b>DaemonSet:</br> запуск демона кластерного хранилища на каждом узле<br> запуск демона сбора журналов на каждом узле</br>запуск демона мониторинга узлов на каждом узле. </br> В простом случае для каждого типа демона будет использоваться один **DaemonSet**, охватывающий все узлы. В более сложной настройке может использоваться несколько наборов<b>DaemonSet</b>для одного типа демона, но с разными флагами и/или разными запросами памяти и процессора для разных типов оборудования. |
| YAML | Язык для сериализации данных c простым синтаксисом, который позволяет хранить сложноорганизованные данные в компактном и читаемом для человека формате |
| JSON | Текстовый формат для сериализации данных, легко читаемый человеком и машиной, основанный на JavaScript |
| Прокси (proxy) | Промежуточный сервер (комплекс программ) в компьютерных сетях, выполняющий роль посредника между пользователем и целевым сервером |
| Агент | Программа, которая вступает в отношение посредничества с пользователем или другой программой |
| Метрика | Собираемые данные, позволяющие получить численное значение потребляемых ресурсов различных целевых объектов (Например, CPU usage, Memory utilization или количество запросов к сервису). Метрики записываются в базе данных Time Series Data Base (TSDB).Проектирование любой системы требует сбора, хранения и отчетности метрик для обеспечения работоспособности системы.Стек для создания отчетов по метрикам в k8s: TSDB , Prometheus/Victoria Metrics и Grafana |
| RBAC (Role-based access control) | система, построенная на использовании ролей и прав пользователей (включая ролевые разрешения, пользовательские роли и контакты между разными ролями). |
| Манифест | Файл в формате YAML или JSON, описывающий любую абстракциюK8s в требуемом формате |
| Метаданные | Метаданные к объектамK8s используются в качестве меток или аннотаций. Метки можно использовать для выбора объектов и для поиска коллекций объектов, которые соответствуют определенным условиям. В отличие от них, аннотации не используются для идентификации и выбора объектов. Метаданные в аннотации могут быть маленькими или большими, структурированными или неструктурированными, кроме этого они могут включать символы, которые не разрешены в метках |
| Плагин | Программа, которая расширяет функциональные возможности других программ |
| API (Application Programming Interface) | Набор компонентов, с помощью которых компьютерная программа может взаимодействовать с другой программой |
| URL-адрес (Uniform Resource Locator) | Cистема унифицированных адресов электронных ресурсов, или единообразный определитель местонахождения ресурса (файла) |
| DNS (Domain Name System) | Компьютерная распределённая система для получения информации о доменах. Чаще всего используется для получения IP-адреса по имени хоста (компьютера или устройства), получения информации о маршрутизации и/или обслуживающих узлах для протоколов в домене |
| UDP (User Datagram Protocol) | Протокол пользовательских дейтограмм, является ненадежным протоколом без установления соединения, не использующим последовательное управление потоком протокола TCP, а предоставляющим свое собственное |
| TCP (Transmission Control Protocol) | Протокол управления передачей, является надежным протоколом с установлением соединений, позволяющим без ошибок доставлять байтовый поток с одной машины на любую другую машину объединенной сети |
| DoT (DNS поверх TLS) | Cтандартный протокол для выполнения разрешения удалённой системы DNS с использованием TLS. Целью этого метода является повышение конфиденциальности и безопасности пользователей путём предотвращения перехвата и манипулирования данными DNS |
| DNSSEC (Domain Name System Security Extensions) Расширения безопасности системы доменных имен | Функция DNS, которая проверяет подлинность ответов на поиск доменных имен. Он не обеспечивает защиту конфиденциальности для этих поисков, но не позволяет злоумышленникам манипулировать или искажать ответы на DNS-запросы |
| Cache (кеш) | Совокупность временных копий файлов программ, а также специально отведенное место их хранения для оперативного доступа |
| Backend | Часть информационной или программной системы, которая работает на сервере и скрыта от конечного пользователя |
| Frontend | Презентационная часть информационной или программной системы, ее пользовательский интерфейс и связанные с ним компоненты |
| DNS64 | Механизм для синтеза записей типа AAAA из записей типа, используется с транслятором IPv6/IPv4 для подключения связи клиент-сервер между клиентом, поддерживающим только IPv6, и сервером, поддерживающим только IPv4 |
| Порт | Число, которое записывается в заголовках протоколов транспортного уровня сетевой модели OSI. Используется для определения программы или процесса-получателя пакета в пределах одного IP-адреса |
| CRD (Custom Resource Definition) | АбстракцияK8s -расширение API Kubernetes, которое не обязательно доступно в установкеK8s по умолчанию. Представляет собой настройку конкретной установкиK8s .Многие основные функцииK8s создаются с использованием CRD, что делаетK8s более модульным |
| Оператор | Программные расширенияK8s , которые используютспециальные ресурсыдля управления приложениями и их компонентами |
| Демон | Программа в UNIX-подобных системах, запускаемая самой системой и работающая в фоновом режиме без прямого взаимодействия с пользователем |
| Аутентификация | Процедура проверки подлинности |
| Авторизация | Предоставление определенному лицу прав на выполнение определенных действий |
| Сертификат | Выпущенный удостоверяющим центром электронный или печатный документ, подтверждающий принадлежность владельцу открытого ключа или каких-либо атрибутов |
| HTTP (Hyper Text Transfer Protocol) | Протокол передачи гипертекстовой разметки, которая используется для передачи данных в интернете |
| HTTPS | Расширение HTTP-протокола, объединение двух протоколов: HTTP и SSL или HTTP и TLS |
| Драйвер | Программный компонент, который позволяет операционной системе и устройству взаимодействовать друг с другом |
| BGP (Border Gateway Protocol) | Протокол, который обеспечивает передачу маршрутов между различными сетями (автономными системами) |
| VPN (Virtual Private Network) | Технология, позволяющая создать безопасное подключение пользователя к сети, организованной между несколькими компьютерами |
| REST API (Representational State Transfer) | Интерфейс, используемый двумя компьютерными системами для обмена информацией через сеть |
| Образ контейнера | Представляет собой двоичные данные, инкапсулирующие приложение и все его программные зависимости.Образы контейнеров — это исполняемые пакеты программного обеспечения, которые могут работать автономно и опираются на среду выполнения. Обычно после создания образ контейнера приложения помещается в реестр, перед обращением к нему в pod |
| Хост | Устройство, предоставляющее сервисы формата "клиент-сервер" в режиме сервера по каким-либо интерфейсам и уникально определённое на этих интерфейсах |
| CI/CD pipeline (continuous integration and continuous deployment) | Концепция непрерывной интеграции и развертывания в процессе разработки программного обеспечения с целью повысить качество продукта |
| Patch | Автоматизированное отдельно поставляемое программное средство, используемое для устранения проблем в программном обеспечении или изменения его функционала, а также сам процесс установки патча |
| Утилита | Вспомогательная компьютерная программа в составе общего программного обеспечения для выполнения специализированных типовых задач, связанных с работой оборудования и операционной системы (ОС) |
| PersistentVolumeClaim (PVC) | Это не что иное как запрос к Persistent Volumes на хранение от пользователя. Это аналог создания pod на ноде. Поды могут запрашивать определенные ресурсы ноды, то же самое делает и PVC. |
| Пространства имен | K8s поддерживает несколько виртуальных кластеров в одном физическом кластере. Такие виртуальные кластеры называются пространствами имён |
| CSR (Certificate Signing Request) | Зашифрованный запрос на выпуск сертификата, содержащий подробную информацию о домене и организации |
| CFSSL | Утилита для работы с PKI/TLS, которая позволяет подписывать, проверять и объединять TLS сертификаты |
| Envsubst | Этоинструмент, который мы можем использовать для динамической генерации конфигурационного файла<b>Nginx</b>перед запуском сервера. |
| StatefulSet | Это объектыK8s , используемые для согласованного развертывания компонентов приложения с отслеживанием состояния. Подам, созданным как часть<b>StatefulSet</b>, присваиваются постоянные идентификаторы, которые они сохраняют даже при изменении расписания.<b>StatefulSet</b>обеспечивает устойчивость и надежность для приложений с состоянием, позволяя им сохранять информацию о состоянии и идентичности даже в случае перезапуска или масштабирования |
| Sidecars | Дополнительные контейнеры вK8s , которые запускаются вместе с основным контейнером приложения |

<a name="id-инструкцияпомиграцииcopenshiftk8s20-маппингресурсов"><b>Маппинг ресурсов</b></a>

Перенос приложения из OpenShift в K8s требует декомпозиции Templates на отдельные YAML-ресурсы, а также замены ряда специфичных для OpenShift объектов на сущности K8s.

| OpenShift| K8s|
| --- | --- |
| DeploymentConfig | Deployment (c особенностями, связанными с BuildConfig) |
| Template | Helm chart |
| Route | Ingress |
| Deployment/Statefulset/Daemonset | Не требуют изменений (только замена параметризации) |
| Service/ConfigMap и т. д. | Не требуют изменений (только замена параметризации) |

[<u><b>Сценарии миграции</b></u>]()

[<u><b>Поддержка Pod Security Admission</b></u>]()

Для реализации требований надежности и безопасности k8s аналогичных настройкам OpenShift на всех кластерах Kubernetes, перед развертыванием приложений необходимо активировать механизм <b>Pod Security Admission</b> (<b>PSA</b>) с политикой <b>restricted</b>. При миграции необходимо следовать указаниям стандартов [<u>STD-7</u>](https://sberworks.ru/wiki/pages/viewpage.action?pageId=37249655 "https://sberworks.ru/wiki/pages/viewpage.action?pageId=37249655") и [<u>STD-22</u>](https://sberworks.ru/wiki/pages/viewpage.action?pageId=132956909 "https://sberworks.ru/wiki/pages/viewpage.action?pageId=132956909").

<b>Pod Security Admission</b> с активированным стандартом <b>restricted</b> запрещает запускать поды в привилегированном и root режиме, запрещает монтировать определенные volumes, использовать linux namespace хоста машины и т.д. Полный список с разрешенными значениями полей представлен в данном сценарии.

Для применения механизма <b>PSA</b> c политикой <b>restricted</b> необходимо в поле labels для манифеста Namespace указать значения:

`pod-security.kubernetes.io/enforce: restricted`

`pod-security.kubernetes.io/enforce-version: v1.25`

`pod-security.kubernetes.io/audit: restricted`

`pod-security.kubernetes.io/audit-version: v1.25`

`pod-security.kubernetes.io/warn: restricted`

`pod-security.kubernetes.io/warn-version: v1.25`

Пример содержания `podsecurity-restricted.yaml`:

```apiVersion: v1

kind: Namespace

metadata:

  name: my-restricted-namespace

  labels:

    pod-security.kubernetes.io/enforce: restricted

    pod-security.kubernetes.io/enforce-version: v1.25

    pod-security.kubernetes.io/audit: restricted

    pod-security.kubernetes.io/audit-version: v1.25

    pod-security.kubernetes.io/warn: restricted

    pod-security.kubernetes.io/warn-version: v1.25

```

В наборе конфигурации `(conf/k8s/base)` для всех рабочих нагрузок, содержащих контейнеры, требуется:

1. Для каждого контейнера (в том числе `initContainers` и `ephemeralContainers`) в составе пода добавить поле `SecurityContext`:

```

kind: Deployment

spec:

  template:

    spec:

      containers|initContainers|ephemeralContainers:

        securityContext:

          capabilities:

            drop:

              - ALL

            privileged: false

            runAsNonRoot: true

            runAsUser: ваш UID компонента

            runAsGroup: ваш GID компонента

            readOnlyRootFilesystem: true

            allowPrivilegeEscalation: false

            seccompProfile:

                type: RuntimeDefault

```

1. Для каждого `Deployment`, `StatefulSet`, `DaemonSet`, `Pod` на уровень `spec` добавить:


```
kind: Deployment

spec:

  template:

    spec:

      securityContext:

        fsGroup: ваш GID владельца тома

seccompProfile:

          type: RuntimeDefault
```

Ниже представлен список того, что запрещает политика `restricted`

Информация представлена в в формате – описание правила; поля в pod/deployment, которые оно затронет; разрешенные значения для этих полей:

1. Использование namespace хостовой машины должно быть запрещено. Поля для этого ограничения:

``````
spec.hostNetwork

spec.hostPID

spec.hostIPC
``````

Разрешенные значения:

`Undefined/nil`

`false`

1. Запрет на запуск контейнеров в привилегированном режиме Поля для этого ограничения:

`spec.containers[*].securityContext.privileged`

`spec.initContainers[*].securityContext.privileged`

`spec.ephemeralContainers[*].securityContext.privileged`

Разрешенные значения:

`Undefined/nil`

`false`

1. Разрешено создавать `Volumes` только со следующими значениями:

Поля для этого ограничения:

`spec.volumes[*]`

Разрешенные значения:

Каждый элемент в spec.volumes[*] должен иметь одно из значений ниже:

`spec.volumes[*].configMap`

`spec.volumes[*].csi`

`spec.volumes[*].downwardAPI`

`spec.volumes[*].emptyDir`

`spec.volumes[*].ephemeral`

`spec.volumes[*].persistentVolumeClaim`

`spec.volumes[*].projected`

`spec.volumes[*].secret`

1. Использование `HostPorts` запрещено.

Поля для этого ограничения:

`spec.containers[*].ports[*].hostPort`

`spec.initContainers[*].ports[*].hostPort`

`spec.ephemeralContainers[*].ports[*].hostPort`

Разрешенные значения:

`Undefined/nil`

`0`

1. Запрещается манипуляция SELinux.

Поля для этого ограничения:

`spec.securityContext.seLinuxOptions.type`

`spec.containers[*].securityContext.seLinuxOptions.type`

`spec.initContainers[*].securityContext.seLinuxOptions.type`

`spec.ephemeralContainers[*].securityContext.seLinuxOptions.type`

Разрешенные значения:

`Undefined/""`

`container_t`

`container_init_t`

`container_kvm_t`

Поля для этого ограничения:

`spec.securityContext.seLinuxOptions.user`

`spec.containers[*].securityContext.seLinuxOptions.user`

`spec.initContainers[*].securityContext.seLinuxOptions.user`

`spec.ephemeralContainers[*].securityContext.seLinuxOptions.user`

`spec.securityContext.seLinuxOptions.role`

`spec.containers[*].securityContext.seLinuxOptions.role`

`spec.initContainers[*].securityContext.seLinuxOptions.role`

`spec.ephemeralContainers[*].securityContext.seLinuxOptions.role`

Разрешенные значения:

`Undefined/""`

1. Тип `Mount /proc`.

Поля для этого ограничения:

`spec.containers[*].securityContext.procMount`

`spec.initContainers[*].securityContext.procMount`

`spec.ephemeralContainers[*].securityContext.procMount`

Разрешенные значения:

`Undefined/nil`

`Default`

1. Параметры `Seccomp` профиля

Поля для этого ограничения:

`spec.securityContext.seccompProfile.type`

`spec.containers[*].securityContext.seccompProfile.type`

`spec.initContainers[*].securityContext.seccompProfile.type`

`spec.ephemeralContainers[*].securityContext.seccompProfile.type`

Разрешенные значения:

`RuntimeDefault`

`Localhost`

1. Настройки `sysctl`.

Поля для этого ограничения:

`spec.securityContext.sysctls[*].name`

Разрешенные значения:

`Undefined/nil`

`kernel.shm_rmid_forced`

`net.ipv4.ip_local_port_range`

`net.ipv4.ip_unprivileged_port_start`

`net.ipv4.tcp_syncookies`

`net.ipv4.ping_group_range`

1. Запрет на `Privilege escalation`.

Поля для этого ограничения:

`spec.containers[*].securityContext.allowPrivilegeEscalation`

`spec.initContainers[*].securityContext.allowPrivilegeEscalation`

`spec.ephemeralContainers[*].securityContext.allowPrivilegeEscalation`

Разрешенные значения:

`false`

1. Запрет на запуск из-под root

Поля для этого ограничения:

`spec.securityContext.runAsNonRoot`

`spec.containers[*].securityContext.runAsNonRoot`

`spec.initContainers[*].securityContext.runAsNonRoot`

`spec.ephemeralContainers[*].securityContext.runAsNonRoot`

Разрешенные значения:

`true`

1. Поле r`unAsUser` не должно быть равно нулю.

Поля для этого ограничения:

`spec.securityContext.runAsUser`

`spec.containers[*].securityContext.runAsUser`

`spec.initContainers[*].securityContext.runAsUser`

`spec.ephemeralContainers[*].securityContext.runAsUser`

Разрешенные значения:

`any non-zero value`

`undefined/null`

1. Контейнеры должны исключить все `Capability`, можно использовать только N`ET_BIND_SERVICE`.

Поля для этого ограничения:

`spec.containers[*].securityContext.capabilities.drop`

`spec.initContainers[*].securityContext.capabilities.drop`

`spec.ephemeralContainers[*].securityContext.capabilities.drop`

Разрешенные значения: Любой список `capabilities` который включает `ALL`

Поля для этого ограничения:

`spec.containers[*].securityContext.capabilities.add`

`spec.initContainers[*].securityContext.capabilities.add`

`spec.ephemeralContainers[*].securityContext.capabilities.add`

Разрешенные значения:

`Undefined/nil`

`NET_BIND_SERVICE`

## Сценарий замены DeploymentConfig на Deployment

В OpenShift часто используется ресурс `DeploymentConfig` это версия Deployment, с расширением за счет ресурсов `ImageStream` и `BuildConfig`: они предназначены для сборки образов и деплоя приложения в кластер.  Прямая замена <b>DeploymentConfig</b> на аналогичный ресурс в k8s невозможна. Поэтому для сохранения функциональности, которую предоставляет связка <b>DeploymentConfig + ImageStream + BuildConfig</b>, потребуются дополнительные инструменты.

Чтобы перенести <b>DeploymentConfig</b> в K8s, можно заменить его на Deployment, а неподдерживаемые функции реализовать сторонними инструментами — например, CI-системой, инструментом для сборки образов, а также внешним <b>registry</b> для их хранения.

1. `apiVersion: apps.openshift.io/v1` замените на `apiVersion: apps/v1 `.
2. `kind: DeploymentConfig` замените на `kind: Deployment `.
3. `spec.selectors` замените с `selector: name: ...` на `selector: matchLabels: name: ...`
4. Убедитесь, что секция `spec.template.spec.containers.image` описана для каждого контейнера.
5. Удалите секции `spec.triggers`, `spec.strategy` и `spec.test`.

Примерный список действий для превращения <b>DeploymentConfig</b> в <b>Deployment</b>:

| OpenShift | K8s |
| --- | --- |
| DeploymentConfig | Deployment ( c особенностями, связанными с ImageStream и BuildConfig) |
| Template | Helm chart |
| Route | Ingress |
| Deployment/Statefulset/Daemonset | Не требуют изменений (только замена параметризации) |
| Service/ConfigMap и т. д. | Не требуют изменений (только замена параметризации) |

<b>Примечание:</b> При миграции из OpenShift в K8s, используйте следующую инструкцию и примеры для переноса <b>Deploment</b> и <b>DeploymentConfig</b>.

[<u><b>Подготовка к миграции</b></u>]()

1. Убедитесь, что установлен и настроен **Kubernetes CLI (kubectl)** для работы с K8s кластером, в который планируется переместить приложение.
2. Убедитесь, что есть доступ к OpenShift кластеру, из которого необходимо перенести приложение, и что имеются права администратора или доступ к проектам.

[<u><b>Сценарий миграции Deployment</b></u>]()

1. Откройте терминал и выполните следующую команду, чтобы выгрузить текущую конфигурацию <b>Deployment</b> из OpenShift:


```oc get deployment <deployment_name> -o yaml --export > deployment.yaml```

1. Внесите необходимые изменения в файл `deployment`.yaml . Если у вас есть специфичные настройки для OpenShift, которых нет в Kubernetes, вам придется изучить и внести соответствующие изменения.
2. Если требуется, удалите все поля, которые не валидны в Kubernetes.
3. Примените новую конфигурацию `Deployment` в Kubernetes с помощью команды:

`kubectl apply -f deployment.yaml`

[<u><b>Сценарий миграции сервисов замены svc -> svc</b></u>]()

1. Создайте манифест (в формате YAML) для сервиса в K8s.

Пример манифеста для svc в K8s:

```apiVersion: v1

kind: Service

metadata:

  name: my-service

spec:

  selector:

    app: my-app

  ports:

    - protocol: TCP

      port: 8080

      targetPort: 8080

  type: LoadBalancer

```

В приведенном примере манифеста сервис с именем `my-service` будет выбирать все pods с `label <app: my-app`, трафик перенаправлен на порт 8080.

Также важно заметить, что тип сервиса установлен на `"LoadBalancer"`, при необходимости, выберите другой тип.

1. Сохраните манифест в файле с расширением <b>.yaml</b> (например, <b>myservice.yaml</b>).
2. Убедитесь, что <b>kubectl</b> настроен для работы с кластером K8s.
3. Запустите миграцию сервиса, выполнив следующую команду в терминале:

`kubectl apply -f myservice.yaml`

Эта команда создаст сервис в K8s кластере на основе манифеста в файле `myservice.yaml`.

1. Проверьте, что сервис успешно создан и работает, выполните команду:

`kubectl get services`

В выводе отобразится сервис со статусом "Running" и внешним IP-адресом (если используется тип LoadBalancer). Сервис успешно мигрирован с OpenShift на k8s.

[<u><b>Сценарий миграции route -> ingress</b></u>]()

1. Создайте объект <b>Ingress</b> в K8s:

```

apiVersion: networking.k8s.io/v1

kind: Ingress

metadata:

  name: example-ingress

spec:

  rules:

  - host: example.com

    http:

      paths:

      - path: /

        pathType: Prefix

        backend:

          service:

            name: example-service

            port:

              number: 8080
```

Где:

`example-ingress` - это имя <b>Ingress</b>;

`example.com` - доменное имя;

`example-service` - имя объекта Service;

`8080` - номер порта сервиса.

1. Установите Ingress Controller (например, Nginx Ingress Controller) в кластере. Этот контроллер будет отслеживать объекты <b>Ingress</b>  и маршрутизировать трафик на соответствующие сервисы.
2. Создайте секрет TLS для доменного имени, если требуется:

`kubectl create secret tls example-tls-secret --cert=example.crt --key=example.key`

Где:

`example-tls-secret` - имя секрета;

`example.crt` - сертификат;

`example.key` - приватный ключ.

1. Обновите объект <b>Ingress</b>, чтобы использовать TLS секрет:

```kind: Ingress

metadata:

  name: example-ingress

spec:

  tls:

  - hosts:

    - example.com

    secretName: example-tls-secret

  rules:

  - host: example.com

    http:

      paths:

      - path: /

        pathType: Prefix

        backend:

          service:

            name: example-service

            port:

              number: 80805.

```

1. Добавьте секцию tls в <b>Ingress</b>, где: `example.com` - доменное имя; `example-tls-secret` - имя TLS секрета.
2. Удалите <b>route</b> на OpenShift:

`oc delete route example-route`

Где `example-route` - имя Route объекта.

После этого маршрут должен успешно мигрировать из OpenShift в Ingress в K8s.

<b>Примечание:</b> Обратите внимание, что инструкция может отличаться в зависимости от вашего специфического сценария и используемых инструментов.

[<u><b>Сценарий миграции daemonset -> daemonset</b></u>]()

1. Создайте манифест для нового <b>DaemonSet</b> K8s.
2. Создайте файл daemonset.yaml и добавьте следующее содержимое:

```kind: DaemonSet

metadata:

  name: my-daemonset

spec:

  selector:

    matchLabels:

      app: my-app

  template:

    metadata:

      labels:

        app: my-app

    spec:

      containers:

      - name: my-container

        image: my-image:latest2.
```
1. Выполните команду для применения манифеста и создания нового <b>DaemonSet</b>:

`kubectl apply -f daemonset.yaml`

1. Cкопируйте Pod с OpenShift на Kubernetes с помощью команды `kubectl cp`.

<b>ПРИМЕР:</b>

`kubectl cp my-daemonset-xyz:/path/to/file /path/to/destination`

1. Удалите <b>DaemonSet</b> в OpenShift с помощью команды:

`oc delete daemonset my-daemonset`

1. Проверьте, что все pods остановлены и удалены.6
2. Выполните команду для запуска нового DaemonSet в K8s:

`kubectl create -f daemonset.yaml`

1. Проверьте, что новый <b>DaemonSet</b> запущен и работает нормально.

[<u><b>Сценарий миграции Secrets</b></u>]()

1. Создайте файл в формате YAML, например `secrets`.`yaml`, и добавьте в него описание Secrets на основе уже существующего Secrets в OpenShift:

```apiVersion: v1

kind: Secret

metadata:

  name: mysecrets

type: Opaque

data:

  username: <OpenShift-username>

  password: <OpenShift-password>
```
1. Замените `OpenShift-username` и `OpenShift-password `на соответствующие значения, закодированные в Base64. Если значение в K8s, вы можете сразу закодировать его в Base64, используя команду в терминале:

`$ echo -n "<OpenShift-username>" | base64`

`$ echo -n "<OpenShift-password>" | base64`

Замените `<OpenShift-username>` и `<OpenShift-password>` на результаты предыдущего шага.

1. Создайте `Secrets` в K8s, используя файл YAML:

`$ kubectl create -f secrets.yaml`

1. Проверьте, что <b>Secrets</b> были успешно созданы:

`$ kubectl get secrets`

В терминале должен отобразиться Secrets с именем `mysecrets`.

1. Для использования <b>Secrets</b> в приложении, сошлитесь на него в манифесте pod или Deployment следующим образом:

```apiVersion: v1

kind: Pod

metadata:

  name: mypod

spec:

  containers:

    - name: mycontainer

      image: myimage

      env:

        - name: USERNAME

          valueFrom:

            secretKeyRef:

              name: mysecrets

              key: username

        - name: PASSWORD

          valueFrom:

            secretKeyRef:

              name: mysecrets

              key: password
```
В приведенном примере переменные среды `USERNAME` и `PASSWORD` будут загружены из `Secrets mysecrets`.

1. Разверните приложение в K8s и убедитесь, что оно успешно получает значения переменных среды из <b>Secrets</b>. <b>Примечание:</b> Убедитесь, что пользователь в K8s имеет необходимые права на создание и управление <b>Secrets</b>. Возможно, понадобится обновить ролевые настройки для пользователя или использовать права администратора для создания <b>Secrets</b>.

[<u><b>Сценарии миграции ConfigMaps</b></u>]()

Перенос <b>ConfigMaps</b> с OpenShift на K8s включает перенос значений, ключей и метаданных.

1. Создайте <b>ConfigMap</b> на OpenShift и сохраните его в YAML-файл, например configmap.yaml . Пример создания <b>ConfigMap</b> с несколькими ключами:

```apiVersion: v1

kind: ConfigMap

metadata:

  name: my-configmap

data:

  key1: value1

  key2: value2

```

1. Создайте новый <b>ConfigMap</b> в K8s с использованием файла configmap.yaml . Можно выполнить команду `kubectl create` или создать конфигурационный файл `kubernetes_configmap.yaml` и применить его с помощью kubectl apply . Например: `kubernetes_configmap.yaml`:

```apiVersion: v1

kind: ConfigMap

metadata:

  name: my-k8s-configmap

data:

  key1: value1

  key2: value2

```

Применение конфигурационного файла:

`kubectl apply -f kubernetes_configmap.yaml`

1. Проверьте, что новый <b>ConfigMap</b> успешно создан, выполнив команду:

`kubectl get configmap my-k8s-configmap`

1. Если существует необходимость восстановить аннотации или метки из `ConfigMap OpenShift` в `ConfigMap K8s`, добавьте их в конфигурационный файл `kubernetes_configmap.yaml`.

Например:

```apiVersion: v1

kind: ConfigMap

metadata:

  name: my-k8s-configmap

  annotations:

    my-annotation: value

  labels:

    my-label: value

data:

  key1: value1

  key2: value2

```

1. Если существует необходимость восстановить аннотации или метки из <b>ConfigMap OpenShift</b> в <b>ConfigMap K8s</b>, добавьте их в конфигурационный файл kubernetes_configmap.yaml .
2. Примените обновленный конфигурационный файл:

`kubectl apply -f kubernetes_configmap.yaml`

1. Проверьте, что обновленный <b>ConfigMap</b> имеет требуемые аннотации и метки, выполнив команду:

`kubectl describe configmap my-k8s-configmap`

При выполнении сценария создан <b>ConfigMap</b> в K8s, который соответствует ConfigMap в OpenShift, включая значения, ключи, аннотации и метки.

[<u><b>Сценарий миграции StatefulSet</b></u>]()

1. Создайте <b>namespace</b> в K8s, где будет развернут <b>StatefulSet</b> и выполните команду:

`kubectl create namespace my-namespace`

1. Переместите манифест <b>StatefulSet</b> из OpenShift в K8s. Откройте манифест <b>StatefulSet</b> из OpenShift и удалите атрибуты и секции, несовместимые с K8s. Некоторые из таких атрибутов могут включать опции маршрутизации, удостоверений и типов сетевых политик. Оставьте только базовую конфигурацию <b>StatefulSet</b>, включая определение шаблона pod и настройку персистентных томов (если они используются).
2. Сохраните изменения, а также сохраните манифест <b>StatefulSet</b> под именем, например, statefulset.yaml .
3. Запустите манифест <b>StatefulSet</b> в K8s, используя команду:

`kubectl apply -f statefulset.yaml -n my-namespace`

Где: statefulset.yaml - путь к сохраненному манифесту <b>StatefulSet</b>; my-namespace - это имя <b>namespace</b>.

<b>StatefulSet</b> будет автоматически развернут и будет управлять pod, как они были настроены в манифесте. Pods будут иметь имена statefulset-name-0 , statefulset-name-1 , ..., где statefulset-name - это имя <b>StatefulSet</b>.

1. Проверьте состояние развернутых pods, используя команду:

`kubectl get pods -n my-namespace`

В выводе должно отобразиться, что pods из <b>StatefulSet</b> запущены и готовы к работе.

При выполнении вышеописанного сценария выполнена миграция <b>StatefulSet</b> с OpenShift на K8s.

[<u><b>Сценарий миграции project -> namespace</b></u>]()

1. Создайте <b>namespace</b> Kubernetes, в котором будет развернут проект.

`kubectl create namespace <namespace_name>`

1. Перенесите манифесты (файлы конфигурации) приложения из OpenSh ift в K8s. Манифесты OpenShift содержат информацию о ресурсах, таких как pods, сервисы и СonfigMaps. Перейдите к директории, содержащей манифесты OpenShift, и адаптируйте их для K8s (например, замените API-ресурсы OpenShift на соответствующие ресурсы K8s).
2. Создайте все ресурсы из манифестов в новом <b>namespace</b> K8s.

`kubectl apply -f <manifest_file> --namespace=<namespace_name>`

1. Убедитесь, что приложение успешно развернуто в новом <b>namespace</b> Kubernetes.

`kubectl get pods --namespace=<namespace_name>`

`kubectl get services --namespace=<namespace_name>`

1. <b>(Необязательный шаг)</b> Если в приложении есть внешний <b>ingress</b> в OpenShift, создайте соответствующий ingress в K8s с использованием средств, предоставляемых кластером или сторонними инструментами.
2. Проверьте, что приложение в новом namespace K8s работает корректно. Проверьте доступность его сервисов и функциональность. Это общие шаги для миграции из OpenShift в K8s. В зависимости от конфигурации вашего проекта и использованных вами функций OpenShift могут потребоваться дополнительные действия и настройки для успешной миграции.

[<u><b>Сценарий миграции Users и Groups</b></u>]()

<b>Рассматриваемые цели:</b>

* Безопасное предоставление сертификата.
* Автоматическая подпись CSR.

<b>Пререквизиты:</b>

* Запущенный кластер K8s.
* Установлен <b>kubectl</b>.
* Имеются привилегированные права администратора.
* Сгенерируйте CSR и закрытый ключ.

1. Используйте CFSSL утилиту для создания конфигурацией в формате JSON.
2. Настройте CFSSL при помощи <b>Envsubst</b> инструмента для генерации закрытого ключа и CSR.
```
{

  "CN": "$USERNAME",

  "names": [

      {

          "O": "$GROUP"

      }

  ],

  "key": {

    "algo": "ecdsa",

    "size": 256

  }

}
```
Kubernetes использует общее имя для сопоставления сертификата с именем пользователя в поле.

Администратор (пользователь) может поделиться этим файлом с новым пользователем и запустить его.

`export USERNAME=**ENTER USERNAME**`

`export GROUP=**ENTER GROUP**`

`envsubst < ../private_key_template.json | cfssl genkey -  | cfssljson -bare client`

Приведенные выше команды сгенерирует два файла `client-key.pem` и `client.csr`.

<b>Примечание:</b> Этот ключевой файл является секретным без права на его распространения.

1. Подпишите CSR и сгенерируйте сертификат для аутентификации.
2. Запросите у нового пользователя прислать вам <b>CSR</b>, <b>имя пользователя</b> и <b>группу</b>, которую он использовал.

<b>Примечание:</b> Доступно использование K8s для подписи CSR и получения клиентского сертификата для аутентификации.

1. С помощью `client.csr` создайте запрос на подписание сертификата `CertificateSigningRequest`.

```export USERNAME=**username*

cat <<EOF | kubectl apply -f -

apiVersion: certificates.k8s.io/v1beta1

kind: CertificateSigningRequest

metadata:

name: $USERNAME

spec:

username: $USERNAME

groups:

- system:authenticated

  request: $(cat client.csr | base64 | tr -d '\n')

  usages:

- digital signature

- key encipherment

- client auth

  EOF

```

По завершению созданного запроса будет создан CSR с запросами на подпись сертификата `kubectl get` `certificatesigningrequests`.

1. Подтвердите CSR и загрузите клиентский сертификат.

`kubectl certificate approve $USERNAME`

`kubectl get csr $USERNAME -o jsonpath={.status.certificate} | base64 --decode > client.pem`

По необходимости отправьте новый файл `client.pem` обратно новому пользователю.

1. Создайте роль и привяжите к ней нового пользователя.
2. Создайте роль с именем пользователя и привязкой роли для предоставления полного доступа к <b>namespace</b>.

```kubectl create namespace $USERNAME

cat <<EOF | kubectl apply -f -

apiVersion: rbac.authorization.k8s.io/v1

kind: Role

metadata:

namespace: default

name: $USERNAME-namespace

rules:

- apiGroups: [""]

  resources: ["*"]

  verbs: ["*"]

  EOF

cat <<EOF | kubectl apply -f -

apiVersion: rbac.authorization.k8s.io/v1

kind: RoleBinding

metadata:

name: access-$USERNAME-namespace

namespace: $USERNAME

subjects:

- kind: User

  name: $USERNAME

  apiGroup: rbac.authorization.k8s.io

  roleRef:

  kind: Role

  name: $USERNAME-namespace

  apiGroup: rbac.authorization.k8s.io

  EOF

```

1. Настройте конфигурацию <b>kubectl</b> для доступа к кластеру.

<b>Примечание:</b> Когда пользователь может пройти аутентификацию и имеет авторизацию в кластере, ему необходимо настроить <b>kubectl</b>.

1. Отправите новому пользователю следующие данные:

* `client.pem`;
* `ca.pem` - полученный из вашего кластера;
* конечную точку API для вашего кластера.

1. После получения данных воспользуйтесь <i>kubectl</i> для изменения конфигурации:

```~/.kube/config

kubectl config set-cluster NAME --server=https://$API_ENDPOINT --certificate-authority=ca.pem --embed-certs=true

kubectl config set-credentials NAME --client-certificate=client.pem --client-key=client-key.pem --embed-certs=true

kubectl config set-context prod --cluster=NAME --namespace=$USERNAME --user=$USERNAME

1. Используйте команду для завершения настройки:

kubectl config use-context NAME

```

[<u><b>Сценарий миграции объектов resourceQuota и LimitRange</b></u>]()

`LimitRange и `ResourceQuota` - это объекты, используемые для управления использованием ресурсов администратором кластера K8s.

[<u><b>ResourceQuota (квота ресурсов)</b></u>]()

<b>ResourceQuota</b> применяет ограничение к процессору (CPU), памяти (RAM) рабочими нагрузками в <b>namespace</b>.

<b>Принцип работы:</b>

Разные команды работают в разных пространствах имен. Реализуется c помощью управление доступом на основе ролей (далее - RBAC)

Администратор создает одну `ResourceQuota` для каждого namespace.

Пользователи создают ресурсы (модули, службы и т.д.) в namespace, а система квот отслеживает использование ресурсов, гарантируя, что они не превысят заданные ограничения, определенных в `ResourceQuota`.

В случае создания или обновления ресурса с превышением установленной квоты, запрос завершится с кодом состояния `Forbidden` с сообщением об ограничении.

Если в пространстве имен включена квота для вычислительных ресурсов CPU и RAM, пользователи должны указывать запросы или ограничения для этих значений; в противном случае система квот может отклонить создание pod. Подсказка: используйте контроллер доступа `LimitRanger` для принудительного использования значений по умолчанию для модулей, которые не требуют вычислительных ресурсов.

[<u><b>Включение ResourceQuota</b></u>]()

Поддержка квот ресурсов включена по умолчанию. Она включается, когда --admission-control= флаг <b>apiserver</b> имеет <b>ResourceQuota</b> в качестве одного из своих аргументов.

<b>ResourceQuota</b> применяется в определенном пространстве имен, когда в этом пространстве имен есть <b>ResourceQuota</b> объект. В пространстве имен должен быть не более одного <b>ResourceQuota</b> объекта.

<b>ResourceQuota - Квота вычислительных ресурсов</b>

Вы можете ограничить общую сумму вычислительных ресурсов, которые могут быть запрошены в данном пространстве имен.

Поддерживаются следующие типы ресурсов:

| <b>Название ресурса</b> | <b>Описание</b> |
| --- | --- |
| cpu | Для всех модулей, находящихся в нетерминальном состоянии, сумма запросов процессора не может превышать это значение |
| limits.cpu | Для всех модулей, находящихся в нетерминальном состоянии, сумма ограничений процессора не может превышать это значение |
| limits.memory | Для всех модулей, находящихся в нетерминальном состоянии, сумма ограничений памяти не может превышать это значение |
| memory | Для всех модулей, находящихся в нетерминальном состоянии, сумма запросов к памяти не может превышать это значение |
| requests.cpu | Для всех модулей, находящихся в нетерминальном состоянии, сумма запросов процессора не может превышать это значение |
| requests.memory | Для всех модулей, находящихся в нетерминальном состоянии, сумма запросов к памяти не может превышать это значение |

[<u><b>Квота ресурсов хранения</b></u>]()

Существует возможность ограничения общего количества ресурсов хранения, которые могут быть запрошены в данном <b>namespace</b>. Кроме того, можно ограничить потребление ресурсов хранилища на основе соответствующего класса хранилища.

| <b>Название ресурса</b> | <b>Описание</b> |
| --- | --- |
| requests.storage | Для всех утверждений о постоянном томе сумма запросов к хранилищу не может превышать это значение |
| persistentvolumeclaims | Общее количество постоянных утверждений о томе, которые могут существовать в пространстве имен |
| .storageclass.storage.k8s.io/requests.storage | По всем постоянным заявкам тома, связанным с именем класса хранилища, сумма запросов к хранилищу не может превышать это значение |
| .storageclass.storage.k8s.io/persistentvolumeclaims | Для всех утверждений о постоянном томе, связанных с именем класса хранилища, общее количество утверждений о постоянном томе, которые могут существовать в namespace |

Например, если оператор (администратор) хочет квотировать хранилище с помощью gold класса хранения, отдельного от класса хранения, оператор (администратор) может определить квоту следующим образом:

* gold.storageclass.storage.k8s.io/requests.storage: 500Gi
* bronze.storageclass.storage.k8s.io/requests.storage: 100Gi

[<u><b>Квота на количество объектов</b></u>]()

Количество объектов заданного типа может быть ограничено. Поддерживаются следующие типы:

| <b>Название ресурса</b> | <b>Описание</b> |
| --- | --- |
| configmaps | Общее количество конфигурационных карт, которые могут существовать в пространстве имен |
| persistentvolumeclaims | Общее количество постоянных утверждений о томе, которые могут существовать в пространстве имен |
| pods | Общее количество модулей в нетерминальном состоянии, которые могут существовать в пространстве имен. Модуль находится в терминальном состоянии, если status.phase in (Failed, Succeeded) имеет значение true |
| replicationcontrollers | Общее количество контроллеров репликации, которые могут существовать в пространстве имен |
| resourcequotas | Общее количество квот ресурсов, которые могут существовать в пространстве имен |
| services | Общее количество сервисов, которые могут существовать в пространстве имен |
| services.loadbalancers | Общее количество служб типа load balancer, которые могут существовать в пространстве имен |
| services.nodeports | Общее количество сервисов типа node port, которые могут существовать в пространстве имен |
| secrets | Общее количество секретов, которые могут существовать в пространстве имен |

Например, <b>квота pods</b> подсчитывается и обеспечивает максимальное количество pods созданных в одном <b>namespace</b>.

Ограничьте создание pods для пространства имен, для избежать случаем создания много маленьких модулей и исчерпывание имеющиеся в кластере IP-адреса pod.

ResourceQuota предназначено для ограничения общего потребления ресурсов <b>namespace</b>, например:

<b>ResourceQuota</b>

```apiVersion: v1

kind: ResourceQuota

metadata:

  name: object-counts

spec:

  hard:

    configmaps: "10"

    persistentvolumeclaims: "4"

    replicationcontrollers: "20"

    secrets: "10"

    services: "10"

```

[<u><b>Сценарий миграции CRD</b></u>]()

CRD мигрирует через CRDProxy - продукт, который позволяет создавать CRD без прав администратора кластера. Ниже представлен альтернативный путь ручной настройки CRD.

[<u><b>Сценарий создания CRD в K8s</b></u>]()

Создайте файл YML с указанием параметра CustomerResourceDefinition

[<u><b>application-manifest.yml</b></u>]()

```

apiVersion: apiextensions.k8s.io/v1beta1

kind: CustomResourceDefinition

metadata:

  name: manifests.cat.kearos.net

spec:

  group: cat.kearos.net

  version: v1

  names:

    kind: ApplicationManifest

    plural: applicationmanifests

    singular: applicationmanifest

    shortNames:

      - cam

  scope: Namespaced

```

Таблица. 1. Инструкция по заполнению YAML файла дляCRD

| <b>№</b> | <b>Наименование функции</b> | <b>Значение</b> |
| --- | --- | --- |
| 1 | <b>apiVersion</b> | Расширение API |
| 2 | <b>kind</b> | Обязательный параметр 'CustomResourceDefinition' дляCRD |
| 3 | <b>name</b> | Имя должно соответствовать приведенным ниже полям спецификации |
| 4 | <b>group</b> | Название группы API, позволяющее сгруппировать несколько ресурсов вместе |
| 5 | <b>names</b> |
| 6 |  | <b>kind</b> | Вид ресурса |
| 7 |  | <b>plural</b> | Название, используемое в Kubernetes API, также используемое по умолчанию для взаимодействия сkubectl |
| 8 |  | <b>singular</b> | Псевдоним для использования API в kubectl и используется в качестве отображаемого значения |
| 9 |  | <b>shortNames</b> | Краткие имена позволяют использовать более короткую строку для соответствия вашему ресурсу в командной строке |
| 10 | <b>scope</b> | Может быть'Namespaced'привязана к определенному пространству имен или'Cluster',где она должна быть уникальной для всего кластера |

[<u><b>Установка CRD</b></u>]()

Используйте приведенный выше пример для создания application-manifest.yml . Для установки CRD в кластер воспользуйтесь командой.

`kubectl create -f application-manifest.yml`

<b>Пример использования ресурсов</b>

Пример:

```apiVersion: cat.kearos.net/v1

kind: ApplicationManifest

metadata:

    name: manifest-cat

spec:

    name: cat

    description: Central Application Tracker

    namespace: cat

    artifactIDs:

        - github.com/joostvdg/cat

    sources:

        - git@github.com:joostvdg/cat.git

```

В параметре `spec` можно поместить любое произвольное поле.

[<u><b>Процесс выполнения миграции использованием инструмента Shifter</b></u>]()

В данном разделе описан облегченный сценарий миграции с использованием инструмента Shifter.

1. Убедитесь что осуществлен доступ к существующему кластеру Open Shift в котором находится приложение для миграции.
2. Выполните oc login с правами администратора:

`oc login -u kubeadmin https://api.exemple:6443`

1. Скачайте и установите Shifter, введите команду:

`curl -OL https:... #актуальная команда - путь`

`cp ./shifter_darwin_amd64 /usr/local/bin/shifter`

`chmod +x /usr/local/bin/shifter`

1. Выполните миграцию:

Получите TOKEN в графической оболочке кластера: и выполните вход с ролью kubeadmin. Затем нажмите на иконку пользователя и выберите пункт <b>Copy login command</b>. На странице отобразится API token:

Рисунок - пункт Copy login command

1. Введите команду:

`OCP_PROJECT=demo-shifter`

`OCP_CLUSTER=https://api.crc.testing:6443`

`TOKEN=shifter cluster -e ${OCP_CLUSTER} -t ${TOKEN} convert -n ${OCP_PROJECT} --output-format yaml`

1. Shifter создаст несколько файлов в директории ./out_test_1/${OCP_PROJECT}. Очистите манифесты от метаполей (<b>creationTimestamp</b>, <b>managedFields</b> и т.п.) а также удалите объекты, относящиеся к кластеру OpenShift (сертификаты, IP адреса сервисов, <b>creationTimestamp</b> и т.п.) и запишите вывод в итоговый манифест для кластера:

<i>Для macOS, Linux или iOS</i>

```brew install yq

cd ./out_test_1/${OCP_PROJECT}

rm -f *ca.crt* default.yaml builder.yaml deployer.yaml

cat *.yaml \

| yq e 'del(.metadata.creationTimestamp)' - \

| yq e 'del(.metadata.resourceVersion)' - \

| yq e 'del(.metadata.selfLink)' - \

| yq e 'del(.metadata.uid)' - \

| yq e 'del(.status)' -  \

| yq e 'del(.metadata.managedFields)' - \

| yq e 'del(.metadata.annotations)' - \

| yq e 'del(.metadata.manager)' - \

| yq e 'del(.metadata.operation)' - \

| yq e 'del(.metadata.generation)' - \

| yq e 'del(.metadata.time)' - \

| yq e 'del(.spec.template.metadata.creationTimestamp)' - \

| yq e 'del(.spec.template.spec.schedulerName)' - \

| yq e 'del(.spec.selector.deploymentconfig)' - \

| yq e 'del(.spec.clusterIP)' - \

| yq e 'del(.spec.clusterIPs)' - \

>> k8s-app-manifest.yaml

```

<i>Для Windows</i>

```winget install yq

cd ./out_test_1/${OCP_PROJECT}

rm -f *ca.crt* default.yaml builder.yaml deployer.yaml

cat *.yaml \

| yq e 'del(.metadata.creationTimestamp)' - \

| yq e 'del(.metadata.resourceVersion)' - \

| yq e 'del(.metadata.selfLink)' - \

| yq e 'del(.metadata.uid)' - \

| yq e 'del(.status)' -  \

| yq e 'del(.metadata.managedFields)' - \

| yq e 'del(.metadata.annotations)' - \

| yq e 'del(.metadata.manager)' - \

| yq e 'del(.metadata.operation)' - \

| yq e 'del(.metadata.generation)' - \

| yq e 'del(.metadata.time)' - \

| yq e 'del(.spec.template.metadata.creationTimestamp)' - \

| yq e 'del(.spec.template.spec.schedulerName)' - \

| yq e 'del(.spec.selector.deploymentconfig)' - \

| yq e 'del(.spec.clusterIP)' - \

| yq e 'del(.spec.clusterIPs)' - \

>> k8s-app-manifest.yaml

1. Экспортируйте <b>openshift project</b> в K8s namespace:

oc get project $OCP_PROJECT -o yaml | \

yq e '.apiVersion |= "v1"' - \

| yq e '.kind |= "Namespace"' - \

| yq e 'del(.metadata.creationTimestamp)' - \

| yq e 'del(.metadata.annotations.*)' - \

| yq e 'del(.metadata.managedFields)' - \

| yq e 'del(.metadata.labels)' - \

| yq e 'del(.metadata.resourceVersion)' - \

| yq e 'del(.metadata.selfLink)' - \

| yq e 'del(.metadata.uid)' - \

| yq e 'del(.status)' - \

> k8s-namespace-manifest.yaml

```

1. Выполните деплоймент мигрируемого приложения в кластер K8s, при помощи команды:

`kubectl apply -f k8s-namespace-manifest.yaml`

`kubectl apply -f k8s-app-manifest.yaml`
