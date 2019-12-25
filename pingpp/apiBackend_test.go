package pingpp

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPingxxClient(t *testing.T) {
	assert.Nil(t, nil)

	appID := "app_aPqnD0bPe1eTePyD"
	Key = "sk_test_autocreateOm5afDPKS8mHTi"
	LogLevel = 2
	AccountPrivateKey = `-----BEGIN RSA PRIVATE KEY-----
MIIEowIBAAKCAQEAuIfYpoep9yG/I2He+LuoCmfeCEkSCqVMf+SlGyHm2SZHSSA2
76kjkxALAUsJm4+GVPUTRYvVGhtmhfv/r7nKPV3dB8GXm6UqpEUQ6MjTirnLVIJR
vGfjzMu7m8vpfTCBcb4k0owRfMDetyR2AI032U2e1cJq5kK0MPDrivX5P8Ek9jYQ
uNxMx/9oYhXvA3xY1D7MKl7kuZMuFZ1QW2IudzzST3WR8aOwAB7EE80CtGeKJK9/
pH8gBUaqFR6V4SogleqY86kc5qhKtGqWsGzbT7pjUAcSjC/hwk0P5AfmP0ysw0h8
f5JjkL9NT9qfNeH8SuZJ31EV9G5hGgJrCWYO3wIDAQABAoIBAHEl0wFhpCVFywau
B2Nq5YHj6HAaJfbjHoI5N8foRL2XzBQZrbXMd2TAbI8RbMW2/r5vO67kK/oDHR1A
xKoCePCh8lrEKOsmHnx1YQXrxlKE1Blwgx/McBDuR44jPXsm1+hfmAKGIl7lgGq5
76Us8jbLcn5N5PBlrElqIHOnmXquMB+UshopZhYk80jzzfe0Epn38NYsHjVg1+U2
emxWYPS/+gZvizomu5izDphvOZiBzJ0dXezzfB2Nh71jqVKuCA7vlbThL87VfXUB
MUGAux1ueO5dinhaJ/8PXc/oz/ce9/CgIyDkqySlFrCUKpMNXYk0vIJq6+o9KYW4
rf89bDkCgYEA3Rw/7DhSAAwqYXqmBLKbqUVgbmXG4XgP51V9G7dK2B84nAVxzIA+
6CcIU84M36ndf0vz+9fYNgGqFztusJMl6gFTfu1YIVSh5P/y4GOdHT1BMnXYMxmY
dMotDeEVjHmXQuZQkm+aGCnjJrZsPKUT/DLZfRPnIy+7B45RwuloYi0CgYEA1aX1
v+fnxeXaGf9Gb+9H4vLtW34xe54e9tNMeg/X+v1QLP+NYiQWZI940R7fTsbPwR72
Hf1nMhItSo7Zj5Sk8yrpZGQv1rq+Wd2HKLQ+XOa+B03/FqfGbn8tlzUspFIb/DJs
G6tn6ZfMBLkTFoADKKOGWGKB6yRNf8nqio2uuLsCgYEA1WCmu4Kad5tF5u5iepQW
riLabY9gf0PLLTcN1gpreueeTMrgSDbW/lh0q9NkDxAxReB5Yueoqm2uxF7Bcjt+
0jVXa10GuJug5IqvyyW2gHoViem5q46e1s4o7oTxs6+/bnsu9DIpJfp+UGSpTVkN
UUuNISxs2UL9ncwVUSDIcD0CgYBoXU+01QYFSAjIlkjEOTD2jhCglv55En6xt/ot
zjuFIeTedl+MR3Cg4uzzyo7vHnMyNXuonragYPy65Rkl6EjDeHzWS5KK5GJD59E0
cwfbveOqEdHAMWjfZUTSDmmmQ23kPoVM5ug11a8Vx4qpfRMN8QOZfVVVlkA9Ylcu
I5PsLQKBgD8g0DbgOJtZpkwNDFa2Lmgyz8R+jq9Vzei3k4svATn8Nernu5+98Ysk
0+jKpU0frxoQmEvB60sYZhgB5HYzSb0Kn4B31bSjtLc1nJeUpkavNSSKUs7lYNhQ
bvcve6Lx/JpQTALxlzHF/DjgXRyG24FBj2/DJPNPx16jNWLFnUeY
-----END RSA PRIVATE KEY-----`
	backend := GetBackend(APIBackend)

	user := &User{}
	err := backend.Call("GET", fmt.Sprintf("/apps/%s/users/%s", appID, "test_business_1"), Key, nil, nil, user)
	assert.Nil(t, err)
	assert.Equal(t, "上海某某科技有限公司2", user.Name)
}
