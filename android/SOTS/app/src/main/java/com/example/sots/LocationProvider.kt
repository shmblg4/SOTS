// LocationProvider.kt
package com.example.sots

import android.Manifest
import android.widget.TextView
import android.content.pm.PackageManager
import android.location.Location
import android.os.Build
import android.os.Looper
import android.telephony.CellInfo
import android.telephony.TelephonyManager
import android.util.Log
import androidx.annotation.RequiresApi
import androidx.appcompat.app.AppCompatActivity
import androidx.core.app.ActivityCompat
import com.google.android.gms.location.*

class LocationProvider(private val activity: AppCompatActivity, private val latTextView:    TextView, private val lonTextView: TextView, private val rsrpTextView: TextView) {

    private lateinit var fusedLocationClient: FusedLocationProviderClient
    private lateinit var locationRequest: LocationRequest
    private lateinit var telephonyManager: TelephonyManager

    init {
        fusedLocationClient = LocationServices.getFusedLocationProviderClient(activity)
        telephonyManager = activity.getSystemService(AppCompatActivity.TELEPHONY_SERVICE) as TelephonyManager
        createLocationRequest()
    }

    private fun createLocationRequest() {
        locationRequest = LocationRequest.create().apply {
            interval = 2000
            fastestInterval = 2000
            priority = LocationRequest.PRIORITY_HIGH_ACCURACY
        }
    }

    fun startLocationUpdates() {
        if (ActivityCompat.checkSelfPermission(activity, Manifest.permission.ACCESS_FINE_LOCATION) != PackageManager.PERMISSION_GRANTED) {
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
        val rsrp = getRSRP() ?: "Unavailable"
        rsrpTextView.text = "RSRP: $rsrp"

        Log.d("LocationInfo", "Lat: ${location.latitude}, Lon: ${location.longitude}, RSRP: $rsrp")
    }

    @RequiresApi(Build.VERSION_CODES.R)
    private fun getRSRP(): String? {
        if (ActivityCompat.checkSelfPermission(activity, Manifest.permission.ACCESS_FINE_LOCATION) != PackageManager.PERMISSION_GRANTED) {
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
}
