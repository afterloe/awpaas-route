package cn.awpaas

import org.springframework.boot.autoconfigure.SpringBootApplication
import org.springframework.boot.runApplication
import org.springframework.cloud.client.discovery.EnableDiscoveryClient

@SpringBootApplication
@EnableDiscoveryClient
class Launch

fun main(args: Array<String>) {
    runApplication<Launch>(*args)
}
