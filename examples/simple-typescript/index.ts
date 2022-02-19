import * as wgpeer from "@spigell/pulumi-wgpeer";

const peer = new wgpeer.New("my-peer", { 
	connection: {
        addr: `${serverIP}:22`,
        user: "rancher",
        key: sshServerPrivateKey
    },
    interface: {
    	privateKey: "aaaa"
    }
});
