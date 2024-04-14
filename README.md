# AWS Challenge : Deploy Web Server and Validate Credit Cards

This repository contains a CloudFormation template (`webServer.json`) that deploys a web server with SSL/TLS support on AWS and golang code (`validateCreditCard.go`) to validate credit card numbers. Below, you'll find an explanation of each component of the template and instructions on how to create a testKeyPair manually through KeyPairs in the AWS console, as well as how to start the web server and access it.

## Template Components

### Parameters
- **KeyPair:** Parameter to specify the name of the keypair used for SSH access.

### Resources

- **VPC:** Creates a Virtual Private Cloud (VPC) with DNS support and hostnames enabled.
- **PublicSubnet:** Creates a subnet within the VPC for hosting public resources.
- **InternetGateway:** Creates an Internet Gateway to enable internet access for resources within the VPC.
- **VPCGatewayAttachment:** Attaches the Internet Gateway to the VPC.
- **PublicRouteTable:** Creates a route table for the public subnet.
- **PublicRoute:** Adds a default route to the Internet Gateway in the public route table.
- **PublicSubnetRouteTableAssociation:** Associates the public subnet with the public route table.
- **PublicSubnetNetworkAclAssociation:** Associates the public subnet with the default network ACL.
- **WebServerSecurityGroup:** Configures security group rules to allow HTTP, HTTPS, and SSH traffic.
- **WebServerInstance:** Launches an EC2 instance with Apache HTTP server installed and configured.
  - Creates a sample index.html file.
  - Generates a self-signed SSL certificate using make-dummy-cert.
  - Configures Apache to use the SSL certificate.
  - Redirects HTTP traffic to HTTPS.
  - Tests the web server configuration

### Outputs
- **URL:** Provides the URL of the sample website hosted on the EC2 instance.

### TestKeyPair Creation (Manually)
1. Navigate to the AWS Management Console.
2. Go to the EC2 dashboard.
3. In the navigation pane, under "Network & Security", select "Key Pairs".
4. Click on the "Create Key Pair" button.
5. Enter a name for the key pair (e.g., "testKeyPair").
6. Click "Create Key Pair".
7. Save the private key file (.pem) to a secure location on your local machine.

### Starting the Webserver and Accessing
1. Deploy the CloudFormation template in your AWS account.
2. During stack creation, provide the necessary parameters, including the KeyPair parameter.
3. Once the stack is created, navigate to the EC2 dashboard.
4. Locate the EC2 instance created by the CloudFormation stack.
5. Connect to the EC2 instance using SSH, providing the private key file created earlier.
   Example command:
   ```bash
   ssh -i "testKeyPair.pem" ec2-user@<public_dns_of_ec2_instance>
