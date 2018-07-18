// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

// Code generated by beats/dev-tools/cmd/asset/asset.go - DO NOT EDIT.

package kubernetes

import (
	"github.com/elastic/beats/libbeat/asset"
)

func init() {
	if err := asset.SetFields("metricbeat", "kubernetes", Asset); err != nil {
		panic(err)
	}
}

// Asset returns asset data
func Asset() string {
	return "eJzsXV9v4zYSf99PMdinLJDmA+ThgF56xQX7p0E22z4cDgEtjW02EqmSlFN/+4KURMkSScs27TiJ9LQrOzM/DofD+Uf6J3jC9TU8lTMUDBXKDwCKqgyv4eNn+/LjB4AUZSJooShn1/CvDwAA7RcgRyVoov9aYIZE4jUsyAcAiUpRtpDX8L+PUmYf/6/fLblQjwlnc7q4hjnJJH4AmFPMUnltCP8EjOTYg6UftS40ZcHLon7jgKWfWzbnIif6NRCWglREUaloIoHPoeCphJwwssAUZusOn6uaQhdNFxEpqESxQmE/cYEKAOvJ7ee7W6gIdkTYPFaUM1Sk874PrgtQ4F8lSnWVZBSZ2vhKg/QJ189cpL3PAnj1c2PoQcopW4BaYsNIBlEIlLwUCcbDcV9RxhSctPsAZDk7JgYf+QGMhBfxAYAhCxdJVkqF4tIwlQVJ8NJK51MQ1wrFLD6s/z483MGA9EBDeelR0IyzxW6cH7giGbAyn6HQC3yUcmZEIUvWV7LMI8GoBSChJn0Jssw1nur/FCVQBjlNBJeYcJaOAxhTUs0cWYR7Cm1WJk/oBsVnf2LS/6h6+RgJNiypVHwhSA4VEDmw1AlnilB2mKVuN4aWXshQL8aaaamIUI+K5m6rkBLV/2CLgL5rgjAgaKVRlE5GfVmM4HRz9wNKSRboEIRv2F0o5m8Hn4YAhahuDJILF+HtxLcx6DJh/fEO2TjUu/tskW/3ubFKp6V+wwXWomeEOU3IAC1hXIvFB3or4JFgK6XAdAtDC4uneFUMjMQmKpmQDNPHecaJ74uVk3cNBYoEmXIr1s7D0AImEkiHrLaP2utR1UbDUwSSZTwhiswy1H8XHG9Gc6pe5YBTnFOGaTUCzd68bY3hhX7jFQrQOZTM/C2mblck44u+ruxtmr7whd5h53xHk0RWhGYa81HM0myt9l9/zYSHiIycayMdO1RISEESqtbaJXFTt3a1/ubbl06lyeMlo03e25eKMezjhUK1JXDxjbHDDz3hTeqHb2VVLNGuE+9wWlhzgWHHIxYqzWgMII9exgdkVMMBqAGSY85F33D49WAy1ODQQKcQT+FQn5dAKjF4h3v+3uXXzgB2dDA9KgCvwcccM+wD3MxaLfyeZldIQh5nYzqXlXL//Xt4nTSAn7l4omwhBzkceFPy+KMaJkhU4+RSkAXOSZkpv554kI9A9M3m2jQb8PCxeyf5k4sT4TG8vKjs6uFczcdHa9t283cRV9xzrmBOM5RrqTDfOcR4Hy6PW0pdJ/y9R2JuCdX+98tFZCeINH44YoyGPa42q5w7Z/gfltitxxp6tpyNChKeZZgo+4laEgVEICyQoSCqKiBX1Q0JomSM9sZLmaSpcXQ+98vZsFOR1x8EeyQdlO+NplJxAYEJF6k0PldbD1I0x+pdQYSiSZkRUYkBlkQCT5JSiN7sh2oehp4ieV/Vdt1AbNRNhVSPNQw2KPAGaikjVO+hAaslYThBy0m/62teF1lGTgRMM9qCq43E5cDx8dd6gzC+VqRqtcHU+usLukLmRSCQSM5iALg3lHblr5nF4P6wLmzcEuaYoyIp2VjWfnXfIvKKEhApeUKN3XmmahkAcdrFmAjUoI6n6oYB5Sws+W7ly1Nb3au14RvJ7ZyHeZr+i7iMDUm9yz8vabKsTfAzke0e5PbW6xaQxxUKSXsr7yBQv1cENwQS7scpaZ/FAex/MPpXiUBTZIrOKQpQvAPE0X9gy+6YzR8zyp4igrn/AgILgVKjqZujfAaBshXPVpg+OjAeyy40PF1yCVkIUtD4mvPz3S2sNrUnMF1PlEVUG81bUxzBOK7xYB3jEWB6vPXaUN5B9HEX7I/bX7bw7uZuD/HnOx07JlM4NetMzTqeJ3azzjetb6+7T2cq27meqWzXe+KV7aa6TA/wVJdxA5/qMoG6DEOl9SaavRZ/v2nlu8cE6cpkbn20bH5ZCC6OvSnf/+3jY7M1b3tCHgRhMqdKnc+cPDjnxKaepyJo9YyU5q9T/XNHAU2lz/YZCOc9VD1bH8DXY9kHdYrm2BbVebTFtnh8rbHWpymZN4Ozj92mufYAj9Tm7N8TtjPYxgRGrnAYmyIZs9Jht1TKbW483t13DRi5c8B7FuOIvQV2MXbvUITuHcgGq7wrsENy2AVPX2UKe4pIq2eKSNvnNU3Iq4tI30XN6EyqJANYZ3reZJfTzO/tBLPeWO3xEtk/XzLu6HLkKtlUEOrBPtd1NZ3jirrYRh/meh1Xu+zdODPciC0u6u7wk4qoMl4qulgS6TdA7gH0BxGy1HY4hhFc1N3yl/BMqDL/UChyykj4CB+S1J8tn3GeIen3To1E2SI0TNzy3WjY0lGQPwdEmcLFhpruCabi48nwBfqru2AOmr8/qhmCC4vqxvTj6km7EUQuv3Be/JskT3w+v4T/CGHi5rsyyy7B/rP+fDi1+tE2oZ59yplmlBcZKkwvW0ncEMa4ui+ZYcHFJfz229fPNMsw/VQP/+pg73jbKqnss88rPLT0u3njj+FScQxMe3Of2ikQCXv3nZvfppRC/vMWXIXARBuCa1CidFmlnaFbMCMFGgI/At+x5H7Szq1qGn2uX3CI2/zGnURQ+wWVW7G18tPM4Mvjbqet8Wx8OcMUi4yv8wOPs3W8mpZgFLcmbkv0ZyfOAY+2+8eRhA5t+Huxr7i4dv1WsYqMJiTahVVuHA2Xfa6ySlHS/qG8LsiDXJJfOlNlC281xxb1hSww8ftv2wv6sTC2dQ7PvHVib3Y6WB1eI4AVqfNcVXRQFZ8hoDNvzI8Z/4Qji4P8Z9Mb3g0q4EK7DZfV9dfa+S3ZE+PPzL9uSiaTJaZlWEkPin8Myg0+IWMY06nu5AC2OLK+jMfY4WlnqptxCLuxTcX5aN61xWRr26fryu/I/KU8pW++BNDYOwlfFnmNNtyX4K4KQ7S5M6m0Y6lmd24K7jjFMpiQo8IxmcR+d8z5lt2p+3j14PV4N1Eju71zMltyqR6Pw1GT9rHdcRPejXG9We5XiDxiPrMHs05o3jcJzTtkKWWLq6urffOYMdEd5nfU3kDAB42J1XJz4b0cou1HZhgrfK4J1idUzjh+7gL1BtBHDFy7/P0RdIxDgvvvHQ8bF9vYSLVAAffVf747DlyNjalfClfYgsRDpa3Hrtj4zPyszbGEVt98Yc6T15xgtjbFxhacKewJnmWO+NgmOMkMQ7YtlhTnZZatG25bpdndW3FeZvHMWkPx/O3aBlKvYXPfO+Oduy3s9ZzZi2bsFTlwgQVPlp9MPft7Dauv/CewtBsSsSq0l7E98vJs9d6uzg2V9wkRXsDqDvKXIYANuNb+HHueO5aOtj9qdl7TbSe5A/Y8prmZ3BHArM01fd6xzG3VNN5pfolhc10NMHCY4W1LVl5TO12DMnima1CGz3QDioPc1Ns5XfYxBDxd9uEGPl320ePWoFnxrMxjlWErYmcYBP5eAfM6ItPtC/Uz3b4w3b7g/sJ0+8Jhg3bdOe+CcoIrDn4d+Ztfp/tttBrMPwEAAP//GmALZw=="
}
