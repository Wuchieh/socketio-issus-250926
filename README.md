# socketio-issus-250926



```go
request.MsgCount.Add(1)

// ignore if response does not contain 'sockets' key
if response.Sockets == nil {
    return
}
request.Sockets.Push(response.Sockets...)

if request.MsgCount.Load() == request.NumSub {
    utils.ClearTimeout(request.Timeout.Load())
    if request.Resolve != nil {
        request.Resolve(types.NewSlice(adapter.SliceMap(request.Sockets.All(), func(client *adapter.SocketResponse) any {
            return socket.SocketDetails(adapter.NewRemoteSocket(client))
        })...))
    }
    r.requests.Delete(requestId)
}
```

Although MsgCount has incremented by +1, the function has been interrupted, so the callback cannot be triggered.