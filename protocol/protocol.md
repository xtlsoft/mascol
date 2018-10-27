# Mascol Protocol

## 0. Basic

Mascol uses tcp to communicate with clients.

## 1. General

Each request will be encoded with json.

For example:

```json
{ "id": 0, "chan": "test", "msg": { "test": "abc" }, "return": false }
```

`id` is the sequence identify, `chan` is the channel name (when it is `_broadcast`, it's a broadcast; when it is `_command`, it's an internal command), `msg` is the message, `return` is whether the request needs return value.

## 2. Broadcasting Message Specifications

The message will be broadcasted to all the clients.

## 3. Commanding Message Specifications

Structure:

(example)

```json
{ "cmd": "subscribe", "param": { "name": "test" } }
```

### 3.1 Subscribe

Params:
| Name | Description | Example |
| ---- | ----------- | ------- |
| name | The name of the channel | test |
