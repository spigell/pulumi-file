import * as file from "@spigell/pulumi-file";
import * as fs from 'fs';

const sshPrivateKey = fs.readFileSync(`../../misc/ssh/id_rsa`, 'utf8');

const peer = new file.Remote("myfile", {
	connection: {
        address: '127.0.0.1:2222',
        user: 'ssh-user',
        privateKey: sshPrivateKey
    },
    useSudo: false,
    writableTempDirectory: '/tmp',
    path: '/config/hello.txt',
    content: 'Greetings from Pulumi!\n'
});
