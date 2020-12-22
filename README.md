# Go SDK to Manage Tasks in the Cloud

Aspose.Tasks for Cloud offers the ability to manipulate and convert Microsoft Project MPT, MPP, MPX & Oracle Primavera XER, XML, and PrimaveraP6XML files in the Cloud. [Aspose.Tasks Cloud SDK for Go](https://products.aspose.cloud/tasks/go) wraps the REST API to make it easier for the developers to integrate Task Management features in their own cloud-based Go applications.

Feel free to explore the [Developer's Guide](https://docs.aspose.cloud/display/taskscloud/Developer+Guide) & [API Reference](https://apireference.aspose.cloud/tasks/) to know all about Aspose.Tasks Cloud API. 

## REST API for Task Management

- [Convert project documents](https://docs.aspose.cloud/tasks/convert-project-document-to-the-specified-format/) to other formats.
- Manipulate task data.
- [Manage project's resources](https://docs.aspose.cloud/tasks/working-with-resources/).
- Handle task links & task assignments.
- Work with project's extended attributes.
- [Read Microsoft Project’s document properties](https://docs.aspose.cloud/tasks/working-with-calendars/) such as start and finish date, tasks scheduling types and so on.
- [Read Microsoft Project’s Calendars](https://docs.aspose.cloud/tasks/working-with-calendars/) and Calendar Exceptions information.

## Read & Write Project Data

**Microsoft Project** MPP, XML, MPT **Primavera** MPX

## Save Project Data As

XER, XLSX, HTML, XML, TXT, TIF, SVG, PNG, JPEG

## Getting Started with Aspose.Tasks Cloud SDK for Go

Firstly, create an account at [Aspose for Cloud](https://dashboard.aspose.cloud/#/apps) to get your application information and free quota to use the API. 

#### Install Aspose.Tasks-Cloud

From Visual Stuio Code:

	Add "github.com/aspose-tasks-cloud/aspose-tasks-cloud-go/api" and "github.com/aspose-tasks-cloud/aspose-tasks-cloud-go/api/models" in the import section of your code

From the command line:

	go get -v github.com/aspose-tasks-cloud/aspose-tasks-cloud-go/api

The complete source code is available at [GitHub Repository](https://github.com/aspose-tasks-cloud/aspose-tasks-cloud-go).

## Sample usage via the SDK

The examples below show how your application have to initiate and get assignments from you project file using Aspose.Tasks-Cloud library:

Config.json file:
```csharp
{
	"AppKey": "your app key",
	"AppSid": "your app sid",
	"BaseUrl": "https://api.aspose.cloud"
} 
```
Go code:

```
	// Start README example

	// init tasks cloud api
	config, _ := models.NewConfiguration(configFilePath)
	tasksApi, ctx, _ := api.CreateTasksApi(config)

	// upload sample.mpp to a cloud
	// uploaded_sample.mpp is a name in the cloud
	file, _ := os.Open("sample.mpp")

	tasksApi.UploadFile(ctx, &requests.UploadFileOpts{
		File: file,
		Path: "uploaded_sample.mpp",
	})

	// get assignments from the project file
	options := &requests.GetAssignmentsOpts{
    		Name:   "uploaded_sample.mpp",
	}

	result, _, err := tasksApi.GetAssignments(ctx, options)

	if err != nil {
		//error handling
	}

	fmt.Println(result.Assignments.AssignmentItem[0].Uid)

	// End README example
```

[Product Page](https://products.aspose.cloud/tasks/go) | [Documentation](https://docs.aspose.cloud/display/taskscloud/Home) | [API Reference](https://apireference.aspose.cloud/tasks/) | [Code Samples](https://github.com/aspose-tasks-cloud/aspose-tasks-cloud-go) | [Blog](https://blog.aspose.cloud/category/tasks/) | [Free Support](https://forum.aspose.cloud/c/tasks) | [Free Trial](https://dashboard.aspose.cloud/#/apps)
