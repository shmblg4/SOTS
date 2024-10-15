package com.example.sots

data class UserData(
    val Login: String,
    val Password: String,
    val Signals: MutableList<Signal> = mutableListOf()
)