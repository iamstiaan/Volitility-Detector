# Shared Params Types

- <a href="https://pkg.go.dev/github.com/iamstiaan/Volitility-Detector/shared">shared</a>.<a href="https://pkg.go.dev/github.com/iamstiaan/Volitility-Detector/shared#OrderParam">OrderParam</a>

# Shared Response Types

- <a href="https://pkg.go.dev/github.com/iamstiaan/Volitility-Detector/shared">shared</a>.<a href="https://pkg.go.dev/github.com/iamstiaan/Volitility-Detector/shared#Order">Order</a>

# Pets

Params Types:

- <a href="https://pkg.go.dev/github.com/iamstiaan/Volitility-Detector">whatsauto</a>.<a href="https://pkg.go.dev/github.com/iamstiaan/Volitility-Detector#PetParam">PetParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/iamstiaan/Volitility-Detector">whatsauto</a>.<a href="https://pkg.go.dev/github.com/iamstiaan/Volitility-Detector#APIResponse">APIResponse</a>
- <a href="https://pkg.go.dev/github.com/iamstiaan/Volitility-Detector">whatsauto</a>.<a href="https://pkg.go.dev/github.com/iamstiaan/Volitility-Detector#Pet">Pet</a>

Methods:

- <code title="post /pet">client.Pets.<a href="https://pkg.go.dev/github.com/iamstiaan/Volitility-Detector#PetService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/iamstiaan/Volitility-Detector">whatsauto</a>.<a href="https://pkg.go.dev/github.com/iamstiaan/Volitility-Detector#PetNewParams">PetNewParams</a>) (<a href="https://pkg.go.dev/github.com/iamstiaan/Volitility-Detector">whatsauto</a>.<a href="https://pkg.go.dev/github.com/iamstiaan/Volitility-Detector#Pet">Pet</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /pet/{petId}">client.Pets.<a href="https://pkg.go.dev/github.com/iamstiaan/Volitility-Detector#PetService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, petID <a href="https://pkg.go.dev/builtin#int64">int64</a>) (<a href="https://pkg.go.dev/github.com/iamstiaan/Volitility-Detector">whatsauto</a>.<a href="https://pkg.go.dev/github.com/iamstiaan/Volitility-Detector#Pet">Pet</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /pet">client.Pets.<a href="https://pkg.go.dev/github.com/iamstiaan/Volitility-Detector#PetService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/iamstiaan/Volitility-Detector">whatsauto</a>.<a href="https://pkg.go.dev/github.com/iamstiaan/Volitility-Detector#PetUpdateParams">PetUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/iamstiaan/Volitility-Detector">whatsauto</a>.<a href="https://pkg.go.dev/github.com/iamstiaan/Volitility-Detector#Pet">Pet</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /pet/{petId}">client.Pets.<a href="https://pkg.go.dev/github.com/iamstiaan/Volitility-Detector#PetService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, petID <a href="https://pkg.go.dev/builtin#int64">int64</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="get /pet/findByStatus">client.Pets.<a href="https://pkg.go.dev/github.com/iamstiaan/Volitility-Detector#PetService.FindByStatus">FindByStatus</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/iamstiaan/Volitility-Detector">whatsauto</a>.<a href="https://pkg.go.dev/github.com/iamstiaan/Volitility-Detector#PetFindByStatusParams">PetFindByStatusParams</a>) ([]<a href="https://pkg.go.dev/github.com/iamstiaan/Volitility-Detector">whatsauto</a>.<a href="https://pkg.go.dev/github.com/iamstiaan/Volitility-Detector#Pet">Pet</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /pet/findByTags">client.Pets.<a href="https://pkg.go.dev/github.com/iamstiaan/Volitility-Detector#PetService.FindByTags">FindByTags</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/iamstiaan/Volitility-Detector">whatsauto</a>.<a href="https://pkg.go.dev/github.com/iamstiaan/Volitility-Detector#PetFindByTagsParams">PetFindByTagsParams</a>) ([]<a href="https://pkg.go.dev/github.com/iamstiaan/Volitility-Detector">whatsauto</a>.<a href="https://pkg.go.dev/github.com/iamstiaan/Volitility-Detector#Pet">Pet</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /pet/{petId}">client.Pets.<a href="https://pkg.go.dev/github.com/iamstiaan/Volitility-Detector#PetService.UpdateByID">UpdateByID</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, petID <a href="https://pkg.go.dev/builtin#int64">int64</a>, body <a href="https://pkg.go.dev/github.com/iamstiaan/Volitility-Detector">whatsauto</a>.<a href="https://pkg.go.dev/github.com/iamstiaan/Volitility-Detector#PetUpdateByIDParams">PetUpdateByIDParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="post /pet/{petId}/uploadImage">client.Pets.<a href="https://pkg.go.dev/github.com/iamstiaan/Volitility-Detector#PetService.UploadImage">UploadImage</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, petID <a href="https://pkg.go.dev/builtin#int64">int64</a>, image <a href="https://pkg.go.dev/builtin#io.Reader">io.Reader</a>, params <a href="https://pkg.go.dev/github.com/iamstiaan/Volitility-Detector">whatsauto</a>.<a href="https://pkg.go.dev/github.com/iamstiaan/Volitility-Detector#PetUploadImageParams">PetUploadImageParams</a>) (<a href="https://pkg.go.dev/github.com/iamstiaan/Volitility-Detector">whatsauto</a>.<a href="https://pkg.go.dev/github.com/iamstiaan/Volitility-Detector#APIResponse">APIResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Store

Response Types:

- <a href="https://pkg.go.dev/github.com/iamstiaan/Volitility-Detector">whatsauto</a>.<a href="https://pkg.go.dev/github.com/iamstiaan/Volitility-Detector#StoreInventoryResponse">StoreInventoryResponse</a>

Methods:

- <code title="post /store/order">client.Store.<a href="https://pkg.go.dev/github.com/iamstiaan/Volitility-Detector#StoreService.NewOrder">NewOrder</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/iamstiaan/Volitility-Detector">whatsauto</a>.<a href="https://pkg.go.dev/github.com/iamstiaan/Volitility-Detector#StoreNewOrderParams">StoreNewOrderParams</a>) (<a href="https://pkg.go.dev/github.com/iamstiaan/Volitility-Detector/shared">shared</a>.<a href="https://pkg.go.dev/github.com/iamstiaan/Volitility-Detector/shared#Order">Order</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /store/inventory">client.Store.<a href="https://pkg.go.dev/github.com/iamstiaan/Volitility-Detector#StoreService.Inventory">Inventory</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/iamstiaan/Volitility-Detector">whatsauto</a>.<a href="https://pkg.go.dev/github.com/iamstiaan/Volitility-Detector#StoreInventoryResponse">StoreInventoryResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Order

Methods:

- <code title="get /store/order/{orderId}">client.Store.Order.<a href="https://pkg.go.dev/github.com/iamstiaan/Volitility-Detector#StoreOrderService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, orderID <a href="https://pkg.go.dev/builtin#int64">int64</a>) (<a href="https://pkg.go.dev/github.com/iamstiaan/Volitility-Detector/shared">shared</a>.<a href="https://pkg.go.dev/github.com/iamstiaan/Volitility-Detector/shared#Order">Order</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /store/order/{orderId}">client.Store.Order.<a href="https://pkg.go.dev/github.com/iamstiaan/Volitility-Detector#StoreOrderService.DeleteOrder">DeleteOrder</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, orderID <a href="https://pkg.go.dev/builtin#int64">int64</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>

# User

Params Types:

- <a href="https://pkg.go.dev/github.com/iamstiaan/Volitility-Detector">whatsauto</a>.<a href="https://pkg.go.dev/github.com/iamstiaan/Volitility-Detector#UserParam">UserParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/iamstiaan/Volitility-Detector">whatsauto</a>.<a href="https://pkg.go.dev/github.com/iamstiaan/Volitility-Detector#User">User</a>

Methods:

- <code title="post /user">client.User.<a href="https://pkg.go.dev/github.com/iamstiaan/Volitility-Detector#UserService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/iamstiaan/Volitility-Detector">whatsauto</a>.<a href="https://pkg.go.dev/github.com/iamstiaan/Volitility-Detector#UserNewParams">UserNewParams</a>) (<a href="https://pkg.go.dev/github.com/iamstiaan/Volitility-Detector">whatsauto</a>.<a href="https://pkg.go.dev/github.com/iamstiaan/Volitility-Detector#User">User</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /user/{username}">client.User.<a href="https://pkg.go.dev/github.com/iamstiaan/Volitility-Detector#UserService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, username <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/iamstiaan/Volitility-Detector">whatsauto</a>.<a href="https://pkg.go.dev/github.com/iamstiaan/Volitility-Detector#User">User</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /user/{username}">client.User.<a href="https://pkg.go.dev/github.com/iamstiaan/Volitility-Detector#UserService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, existingUsername <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/iamstiaan/Volitility-Detector">whatsauto</a>.<a href="https://pkg.go.dev/github.com/iamstiaan/Volitility-Detector#UserUpdateParams">UserUpdateParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="delete /user/{username}">client.User.<a href="https://pkg.go.dev/github.com/iamstiaan/Volitility-Detector#UserService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, username <a href="https://pkg.go.dev/builtin#string">string</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="post /user/createWithList">client.User.<a href="https://pkg.go.dev/github.com/iamstiaan/Volitility-Detector#UserService.NewWithList">NewWithList</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/iamstiaan/Volitility-Detector">whatsauto</a>.<a href="https://pkg.go.dev/github.com/iamstiaan/Volitility-Detector#UserNewWithListParams">UserNewWithListParams</a>) (<a href="https://pkg.go.dev/github.com/iamstiaan/Volitility-Detector">whatsauto</a>.<a href="https://pkg.go.dev/github.com/iamstiaan/Volitility-Detector#User">User</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /user/login">client.User.<a href="https://pkg.go.dev/github.com/iamstiaan/Volitility-Detector#UserService.Login">Login</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/iamstiaan/Volitility-Detector">whatsauto</a>.<a href="https://pkg.go.dev/github.com/iamstiaan/Volitility-Detector#UserLoginParams">UserLoginParams</a>) (<a href="https://pkg.go.dev/builtin#string">string</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /user/logout">client.User.<a href="https://pkg.go.dev/github.com/iamstiaan/Volitility-Detector#UserService.Logout">Logout</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
