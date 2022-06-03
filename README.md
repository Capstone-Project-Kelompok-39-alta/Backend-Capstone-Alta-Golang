# Backend-Capstone-Alta-Golang



Struktur Folder
- APP = Isinya hanya main.go
- Infrastructure = isi nya Folder Database(database dan env) dan Http(route, middleware, server)
- Interface/Controller = isinya file controller/layer Controller
- Repository = isinya file repository/layer repository
- Service = isinya file service/layer service
- Entities = isinya file entities/layer entities(model)
- domains = isinya file konfigurasi dari layer repository dan service 

NB : pada layer repository, service, dan entities kita sesuaikan dengan apa yang kita kerjakan a.k.a admin ataupun user
```/admin/repository/login/nama file.go```