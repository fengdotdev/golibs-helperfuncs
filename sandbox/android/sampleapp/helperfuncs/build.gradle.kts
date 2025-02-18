plugins {
    id("java-library")
    alias(libs.plugins.jetbrains.kotlin.jvm)

}
java {
    sourceCompatibility = JavaVersion.VERSION_11
    targetCompatibility = JavaVersion.VERSION_11
}
kotlin {
    compilerOptions {
        jvmTarget = org.jetbrains.kotlin.gradle.dsl.JvmTarget.JVM_11
    }
}


dependencies {
    implementation(libs.androidx.annotation.jvm)
    testImplementation(kotlin("test"))
    testImplementation("junit:junit:4.13.2")
}
