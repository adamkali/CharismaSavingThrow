package controller

type InputComponent struct {
    Name string
    Type string
    Placeholder string
    Value string
    Disabled bool
}

type ButtonComponent struct {
    Name string
    HtmxMethod string
    HtmxEndpoint string
    HtmxTarget string
    Text string
    Icon string
}

type IconButtonComponent struct {
    Name string
    HtmxMethod string
    HtmxEndpoint string
    HtmxTarget string
    Text string
    Icon string
}

type SubmitButtonComponent struct {
    Name string
    Text string
    Icon string
}
