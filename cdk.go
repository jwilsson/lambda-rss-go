package main

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type StackProps struct {
	awscdk.StackProps
}

func NewLambdaRssStack(scope constructs.Construct, id string, props *StackProps) awscdk.Stack {
	var sprops awscdk.StackProps
	if props != nil {
		sprops = props.StackProps
	}

	stack := awscdk.NewStack(scope, &id, &sprops)
	function := awslambda.NewFunction(stack, jsii.String("LambdaRssFunction"), &awslambda.FunctionProps{
		Architecture: awslambda.Architecture_ARM_64(),
		Runtime:      awslambda.Runtime_PROVIDED_AL2023(),
		Handler:      jsii.String("bootstrap"),
		Code:         awslambda.Code_FromAsset(jsii.String("./app/build/"), nil),
	})

	functionUrl := function.AddFunctionUrl(&awslambda.FunctionUrlOptions{
		AuthType: awslambda.FunctionUrlAuthType_NONE,
	})

	function.AddPermission(jsii.String("LambdaRssInvokeFunctionPermission"), &awslambda.Permission{
		Action:    jsii.String("lambda:InvokeFunction"),
		Principal: awsiam.NewServicePrincipal(jsii.String("lambda.amazonaws.com"), nil),
	});

	awscdk.NewCfnOutput(stack, jsii.String("lambdaRssFunctionUrlOutput"), &awscdk.CfnOutputProps{
		Value: functionUrl.Url(),
	})

	return stack
}

func main() {
	defer jsii.Close()

	app := awscdk.NewApp(nil)

	NewLambdaRssStack(app, "LambdaRssStack", &StackProps{})

	app.Synth(nil)
}
