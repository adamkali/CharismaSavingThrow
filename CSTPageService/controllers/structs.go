package controller

type InputComponent struct {
    Label string
    Disabled bool
    Value string
    Name string
    Type string
}

type ButtonComponent struct {
    Endpoint string 
    Target string
    Action string
    Text string
    SvgName string
    Method string
}

type IconButtonComponent struct {
    Endpoint string 
    Target string
    Action string
    Text string
    SvgName string
    Method string
}

type SubmitButtonComponent struct { 
    Text string
    SvgName string
    
}
