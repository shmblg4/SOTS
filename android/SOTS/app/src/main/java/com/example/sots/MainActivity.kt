package com.example.sots

import android.Manifest
import android.content.pm.PackageManager
import android.location.Location
import android.os.Build
import android.os.Bundle
import android.os.Looper
import android.telephony.TelephonyManager
import android.telephony.CellInfo
import android.util.Log
import android.widget.TextView
import android.widget.Toast
import androidx.annotation.RequiresApi
import androidx.appcompat.app.AppCompatActivity
import androidx.core.app.ActivityCompat
import com.google.android.gms.location.*

class MainActivity : AppCompatActivity() {

    private lateinit var fusedLocationClient: FusedLocationProviderClient
    private lateinit var locationRequest: LocationRequest
    private lateinit var latTextView: TextView
    private lateinit var lonTextView: TextView
    private lateinit var rsrpTextView: TextView
    private lateinit var telephonyManager: TelephonyManager

    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)
        setContentView(R.layout.activity_main)
        latTextView = findViewById(R.id.latTextView)
        lonTextView = findViewById(R.id.lonTextView)
        rsrpTextView = findViewById(R.id.rsrpTextView)

        fusedLocationClient = LocationServices.getFusedLocationProviderClient(this)
        telephonyManager = getSystemService(TELEPHONY_SERVICE) as TelephonyManager

        if (ActivityCompat.checkSelfPermission(this, Manifest.permission.ACCESS_FINE_LOCATION) != PackageManager.PERMISSION_GRANTED) {
            ActivityCompat.requestPermissions(this, arrayOf(Manifest.permission.ACCESS_FINE_LOCATION), 1)
        } else {
            createLocationRequest()
            startLocationUpdates()
        }
    }

    private fun createLocationRequest() {
        locationRequest = LocationRequest.create().apply {
            interval = 2000
            fastestInterval = 2000
            priority = LocationRequest.PRIORITY_HIGH_ACCURACY
        }
    }

    private fun startLocationUpdates() {
        if (ActivityCompat.checkSelfPermission(this, Manifest.permission.ACCESS_FINE_LOCATION) != PackageManager.PERMISSION_GRANTED) {
            return
        }

        fusedLocationClient.requestLocationUpdates(locationRequest, object : LocationCallback() {
            @RequiresApi(Build.VERSION_CODES.R)
            override fun onLocationResult(locationResult: LocationResult) {
                super.onLocationResult(locationResult)

                for (location in locationResult.locations) {
                    updateUI(location)
                }
            }
        }, Looper.getMainLooper())
    }

    @RequiresApi(Build.VERSION_CODES.R)
    private fun updateUI(location: Location) {
        latTextView.text = "Latitude: ${location.latitude}"
        lonTextView.text = "Longitude: ${location.longitude}"
        rsrpTextView.text = "RSRP: ${getRSRP() ?: "Unavailable"}"

        Log.d("LocationInfo", "Lat: ${location.latitude}, Lon: ${location.longitude}, RSRP: ${rsrpTextView.text}")
    }

    @RequiresApi(Build.VERSION_CODES.R)
    private fun getRSRP(): String? {
        if (ActivityCompat.checkSelfPermission(this, Manifest.permission.ACCESS_FINE_LOCATION) != PackageManager.PERMISSION_GRANTED) {
            return null
        }

        val cellInfoList: List<CellInfo> = telephonyManager.allCellInfo
        for (cellInfo in cellInfoList) {
            if (cellInfo.isRegistered) {
                val signalStrength = cellInfo.cellSignalStrength
                return "${signalStrength.dbm} dBm"
            }
        }
        return null
    }
    override fun onRequestPermissionsResult(requestCode: Int, permissions: Array<out String>, grantResults: IntArray) {
        super.onRequestPermissionsResult(requestCode, permissions, grantResults)
        if (requestCode == 1) {
            if ((grantResults.isNotEmpty() && grantResults[0] == PackageManager.PERMISSION_GRANTED)) {
                startLocationUpdates()
            } else {
                Toast.makeText(this, "Permission denied. Location features will be limited.", Toast.LENGTH_SHORT).show()
            }
        }
    }
}