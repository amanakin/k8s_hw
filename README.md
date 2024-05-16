# k8s_hw

## Задание 1

Запуск:
```shell
minikube start
eval $(minikube docker-env)
./scripts/deploy.sh
```

Проверить, что все работает и получить адрес сервера:
```shell
kubectl get pods
minikube service timer-app --url
```

Чтобы проверить API можно сделать запросы к серверу по полученному IP (через прокси minikube):
```shell
curl -v http://127.0.0.1:55496/time
```

Проверить скрапер статистики можно на дашборде кубера: 
![img.png](imgs/dashboard_scraper.png)

## Задание 2

Установка `istioctl`. Если у вас MacOS, запустите следующую команду: 
```shell
./scripts/setup.sh
```

Запуск `minikube`
```shell
minikube start
```

Деплой приложений:
```shell
./scripts/deploy.sh
```

После деплоя можно проверить работу на любом интерфейсе, к примеру localhost:
```
amanakin@amanakin-osx:~$ curl -v localhost/time
*   Trying [::1]:80...
* Connected to localhost (::1) port 80
> GET /time HTTP/1.1
> Host: localhost
> User-Agent: curl/8.4.0
> Accept: */*
>
< HTTP/1.1 200 OK
< content-type: application/json
< date: Thu, 16 May 2024 08:17:07 GMT
< content-length: 37
< x-envoy-upstream-service-time: 217
< server: istio-envoy
<
{"time":"2024-05-16T11:17:07+03:00"}
* Connection #0 to host localhost left intact
```

Вы сделаете запрос к `istio` и через `ingress-gateway` он попадет в приложение.

Если зайти на под `scraper-app`, то можно увидеть статистику запросов:
```
root@scraper-app-68756b7cbb-pckkf:/app# tail -f statistics_log.txt 
{"requests": 4}
{"requests": 4}
{"requests": 4}
{"requests": 4}
{"requests": 4}
{"requests": 4}
{"requests": 4}
{"requests": 4}
{"requests": 4}
{"requests": 4}
{"requests": 4}
{"requests": 4}
{"requests": 4}
{"requests": 4}
{"requests": 5}
```