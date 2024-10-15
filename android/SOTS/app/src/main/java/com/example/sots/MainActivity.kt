package com.example.sots

import android.Manifest
import android.content.pm.PackageManager
import android.os.Build
import android.os.Bundle
import android.widget.TextView
import android.widget.Toast
import androidx.annotation.RequiresApi
import androidx.appcompat.app.AppCompatActivity
import androidx.core.app.ActivityCompat

class MainActivity : AppCompatActivity() {

    private lateinit var latTextView: TextView
    private lateinit var lonTextView: TextView
    private lateinit var rsrpTextView: TextView
    private lateinit var locationProvider: LocationProvider

    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)
        setContentView(R.layout.activity_main)
        latTextView = findViewById(R.id.latTextView)
        lonTextView = findViewById(R.id.lonTextView)
        rsrpTextView = findViewById(R.id.rsrpTextView)

        locationProvider = LocationProvider(this, latTextView, lonTextView, rsrpTextView)

        if (ActivityCompat.checkSelfPermission(this, Manifest.permission.ACCESS_FINE_LOCATION) != PackageManager.PERMISSION_GRANTED) {
            ActivityCompat.requestPermissions(this, arrayOf(Manifest.permission.ACCESS_FINE_LOCATION), 1)
        } else {
            locationProvider.startLocationUpdates()
        }
    }

    override fun onRequestPermissionsResult(requestCode: Int, permissions: Array<out String>, grantResults: IntArray) {
        super.onRequestPermissionsResult(requestCode, permissions, grantResults)
        if (requestCode == 1) {
            if ((grantResults.isNotEmpty() && grantResults[0] == PackageManager.PERMISSION_GRANTED)) {
                locationProvider.startLocationUpdates()
            } else {
                Toast.makeText(this, "Permission denied. Location features will be limited.", Toast.LENGTH_SHORT).show()
            }
        }
    }
}
