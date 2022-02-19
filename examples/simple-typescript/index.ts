import * as file from "@spigell/pulumi-file";

const peer = new file.NewRemote("myfile", { 
	connection: {
        addr: '127.0.0.1:2222',
        user: 'ssh-user',
        key: sshServerPrivateKey
    },
    path: '/config/hello.txt',
    content: 'Greetings from Pulumi!'
});
