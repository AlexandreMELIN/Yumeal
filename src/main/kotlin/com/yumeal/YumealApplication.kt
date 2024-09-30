package com.yumeal

import org.springframework.boot.autoconfigure.SpringBootApplication
import org.springframework.boot.runApplication

@SpringBootApplication
class YumealApplication

fun main(args: Array<String>) {
    runApplication<YumealApplication>(*args)
}
