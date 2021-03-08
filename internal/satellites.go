package quasar


type DataInterpreter interface {
    GetMessage(...[]string)(msg string)
    GetLocation(distances ...float32)(x,y float32)
}