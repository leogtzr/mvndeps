# mvndeps

Little CLI program to show dependencies from a pom.xml mvn file.

```
[0 [00:56][leo@mvndeps]$ ./mvndeps pom.xml 
[groupId: org.springframework.boot], [artifactId: spring-boot-starter-batch], [version: ]
[groupId: org.springframework.boot], [artifactId: spring-boot-starter-test], [version: ]
[groupId: org.springframework.batch], [artifactId: spring-batch-test], [version: ]
[groupId: org.springframework.boot], [artifactId: spring-boot-starter-jdbc], [version: ]
[groupId: mysql], [artifactId: mysql-connector-java], [version: ]
[groupId: com.thoughtworks.xstream], [artifactId: xstream], [version: 1.4.7]
[groupId: org.springframework], [artifactId: spring-oxm], [version: 4.1.2.RELEASE]
[0 [00:56][leo@mvndeps]$ 
```