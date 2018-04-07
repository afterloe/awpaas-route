package cn.awpaas

import org.springframework.boot.autoconfigure.SpringBootApplication
import org.springframework.boot.runApplication
import org.springframework.cloud.client.discovery.EnableDiscoveryClient
import java.io.Serializable

@SpringBootApplication
@EnableDiscoveryClient
class Launch : Serializable

fun main(args: Array<String>) {
    runApplication<Launch>(*args)
}
