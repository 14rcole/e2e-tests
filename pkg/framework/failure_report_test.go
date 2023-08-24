package framework

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

const plainLogs = `
2023-08-18T01:18:56.213Z	INFO	ComponentImageRepository	controllers/component_image_controller.go:249	Prepared image repository build-e2e-rsql-tenant/test-app-ngqh/build-suite-test-component-image-source-ajdr for Component	{"controller": "component", "controllerGroup": "appstudio.redhat.com", "controllerKind": "Component", "Component": {"name":"build-suite-test-component-image-source-ajdr","namespace":"build-e2e-rsql-tenant"}, "namespace": "build-e2e-rsql-tenant", "name": "build-suite-test-component-image-source-ajdr", "reconcileID": "1c1f7548-b16a-43ba-a91b-0a2aa32cc6cd", "action": "ADD"}
2023-08-18T01:18:57.532Z	INFO	ComponentImageRepository	controllers/component_image_controller.go:259	Prepared image registry push secret redhat-appstudio-qe+build-e2e-rsql-tenanttest-app-ngqhbuild-suite-test-component-image-source-ajdr for Component	{"controller": "component", "controllerGroup": "appstudio.redhat.com", "controllerKind": "Component", "Component": {"name":"build-suite-test-component-image-source-ajdr","namespace":"build-e2e-rsql-tenant"}, "namespace": "build-e2e-rsql-tenant", "name": "build-suite-test-component-image-source-ajdr", "reconcileID": "1c1f7548-b16a-43ba-a91b-0a2aa32cc6cd", "action": "UPDATE"}
2023-08-18T01:18:58.654Z	INFO	ComponentImageRepository	controllers/component_image_controller.go:269	Prepared remote secret build-suite-test-component-image-source-ajdr-pull for Component	{"controller": "component", "controllerGroup": "appstudio.redhat.com", "controllerKind": "Component", "Component": {"name":"build-suite-test-component-image-source-ajdr","namespace":"build-e2e-rsql-tenant"}, "namespace": "build-e2e-rsql-tenant", "name": "build-suite-test-component-image-source-ajdr", "reconcileID": "1c1f7548-b16a-43ba-a91b-0a2aa32cc6cd", "action": "UPDATE"}
2023-08-18T01:18:59.654Z	INFO	ComponentImageRepository	controllers/component_image_controller.go:293	Image repository finalizer added to the Component update	{"controller": "component", "controllerGroup": "appstudio.redhat.com", "controllerKind": "Component", "Component": {"name":"build-suite-test-component-image-source-ajdr","namespace":"build-e2e-rsql-tenant"}, "namespace": "build-e2e-rsql-tenant", "name": "build-suite-test-component-image-source-ajdr", "reconcileID": "1c1f7548-b16a-43ba-a91b-0a2aa32cc6cd", "action": "UPDATE"}
2023-08-18T01:19:03.673Z	INFO	ComponentImageRepository	controllers/component_image_controller.go:299	Component updated successfully	{"controller": "component", "controllerGroup": "appstudio.redhat.com", "controllerKind": "Component", "Component": {"name":"build-suite-test-component-image-source-ajdr","namespace":"build-e2e-rsql-tenant"}, "namespace": "build-e2e-rsql-tenant", "name": "build-suite-test-component-image-source-ajdr", "reconcileID": "1c1f7548-b16a-43ba-a91b-0a2aa32cc6cd", "action": "UPDATE"}
2023-08-18T01:19:57.257Z	INFO	ComponentImageRepository	controllers/component_image_controller.go:170	Waiting for devfile model in component	{"controller": "component", "controllerGroup": "appstudio.redhat.com", "controllerKind": "Component", "Component": {"name":"devfile-sample-hello-world-0dwf","namespace":"build-e2e-bslz-tenant"}, "namespace": "build-e2e-bslz-tenant", "name": "devfile-sample-hello-world-0dwf", "reconcileID": "188862ec-a820-473e-a3c1-a2bb87031138"}
2023-08-18T01:19:57.327Z	INFO	ComponentImageRepository	controllers/component_image_controller.go:170	Waiting for devfile model in component	{"controller": "component", "controllerGroup": "appstudio.redhat.com", "controllerKind": "Component", "Component": {"name":"devfile-sample-hello-world-0dwf","namespace":"build-e2e-bslz-tenant"}, "namespace": "build-e2e-bslz-tenant", "name": "devfile-sample-hello-world-0dwf", "reconcileID": "88549a15-0c83-4ac0-9e2d-f754e4aea532"}`

const jsonLogs = `
{"level":"info","ts":"2023-08-17T04:58:49Z","caller":"controller/controller.go:228","msg":"Starting workers","controller":"systemconfig","controllerGroup":"jvmbuildservice.io","controllerKind":"SystemConfig","worker count":1}
{"level":"info","ts":"2023-08-18T01:28:06Z","logger":"artifactbuild","caller":"artifactbuild/artifactbuild.go:234","msg":"ArtifactBuild hacbs.test.simple.gradle.jdk8.1.1-f6a4dce0 changing state from ArtifactBuildDiscovering to ArtifactBuildBuilding","namespace":"rhtap-demo-afcg-tenant","resource":"hacbs.test.simple.gradle.jdk8.1.1-f6a4dce0","kind":"ArtifactBuild","ab-gav":"io.github.stuartwdouglas.hacbs-test.gradle:hacbs-test-simple-gradle-jdk8:1.1","ab-initial-state":"ArtifactBuildDiscovering"}
{"level":"info","ts":"2023-08-18T01:29:06Z","logger":"artifactbuild","caller":"artifactbuild/artifactbuild.go:575","msg":"Updating label from  to building to match ArtifactBuildDiscovering","namespace":"rhtap-demo-afcg-tenant","resource":"shaded.jdk11.1.9-c65abf6b"}
{"level":"info","ts":"2023-08-18T01:30:06Z","logger":"artifactbuild","caller":"artifactbuild/artifactbuild.go:234","msg":"ArtifactBuild shaded.jdk11.1.9-c65abf6b changing state from ArtifactBuildDiscovering to ArtifactBuildBuilding","namespace":"rhtap-demo-afcg-tenant","resource":"shaded.jdk11.1.9-c65abf6b","kind":"ArtifactBuild","ab-gav":"io.github.stuartwdouglas.hacbs-test.shaded:shaded-jdk11:1.9","ab-initial-state":"ArtifactBuildDiscovering"}
{"level":"info","ts":"2023-08-18T01:31:06Z","logger":"artifactbuild","caller":"artifactbuild/artifactbuild.go:234","msg":"ArtifactBuild shaded.jdk11.1.9-c65abf6b changing state from ArtifactBuildBuilding to ArtifactBuildNew","namespace":"rhtap-demo-afcg-tenant","resource":"shaded.jdk11.1.9-c65abf6b","kind":"ArtifactBuild","ab-gav":"io.github.stuartwdouglas.hacbs-test.shaded:shaded-jdk11:1.9","ab-initial-state":"ArtifactBuildBuilding"}
{"level":"info","ts":"2023-08-18T01:32:06Z","logger":"artifactbuild","caller":"artifactbuild/artifactbuild.go:524","msg":"Found pipeline run with community dependencies","namespace":"rhtap-demo-afcg-tenant","resource":"hacbs-test-project-jyxg-on-push-vxwtr","kind":"PipelineRun"}
{"level":"info","ts":"2023-08-18T01:33:06Z","logger":"artifactbuild","caller":"artifactbuild/artifactbuild.go:530","msg":"Found community dependency, creating ArtifactBuild","namespace":"rhtap-demo-afcg-tenant","resource":"hacbs-test-project-jyxg-on-push-vxwtr","kind":"PipelineRun","gav":"io.github.stuartwdouglas.hacbs-test.gradle:hacbs-test-simple-gradle-jdk8:1.1","artifactbuild":"hacbs.test.simple.gradle.jdk8.1.1-f6a4dce0","action":"ADD"}
{"level":"info","ts":"2023-08-18T01:34:06Z","logger":"artifactbuild","caller":"artifactbuild/artifactbuild.go:530","msg":"Found community dependency, creating ArtifactBuild","namespace":"rhtap-demo-afcg-tenant","resource":"hacbs-test-project-jyxg-on-push-vxwtr","kind":"PipelineRun","gav":"io.github.stuartwdouglas.hacbs-test.shaded:shaded-jdk11:1.9","artifactbuild":"shaded.jdk11.1.9-c65abf6b","action":"ADD"}
{"level":"info","ts":"2023-08-18T01:35:06Z","logger":"artifactbuild","caller":"artifactbuild/artifactbuild.go:530","msg":"Found community dependency, creating ArtifactBuild","namespace":"rhtap-demo-afcg-tenant","resource":"hacbs-test-project-jyxg-on-push-vxwtr","kind":"PipelineRun","gav":"io.github.stuartwdouglas.hacbs-test.simple:simple-jdk17:0.1.2","artifactbuild":"simple.jdk17.0.1.2-22fafbfd","action":"ADD"}`

func TestPlainLogParsing(t *testing.T) {

	start, _ := time.Parse(time.RFC3339, "2023-08-18T01:18:58.100Z")
	filtered := FilterLogs(plainLogs, start)
	assert.Equal(t,
		`2023-08-18T01:18:58.654Z	INFO	ComponentImageRepository	controllers/component_image_controller.go:269	Prepared remote secret build-suite-test-component-image-source-ajdr-pull for Component	{"controller": "component", "controllerGroup": "appstudio.redhat.com", "controllerKind": "Component", "Component": {"name":"build-suite-test-component-image-source-ajdr","namespace":"build-e2e-rsql-tenant"}, "namespace": "build-e2e-rsql-tenant", "name": "build-suite-test-component-image-source-ajdr", "reconcileID": "1c1f7548-b16a-43ba-a91b-0a2aa32cc6cd", "action": "UPDATE"}
2023-08-18T01:18:59.654Z	INFO	ComponentImageRepository	controllers/component_image_controller.go:293	Image repository finalizer added to the Component update	{"controller": "component", "controllerGroup": "appstudio.redhat.com", "controllerKind": "Component", "Component": {"name":"build-suite-test-component-image-source-ajdr","namespace":"build-e2e-rsql-tenant"}, "namespace": "build-e2e-rsql-tenant", "name": "build-suite-test-component-image-source-ajdr", "reconcileID": "1c1f7548-b16a-43ba-a91b-0a2aa32cc6cd", "action": "UPDATE"}
2023-08-18T01:19:03.673Z	INFO	ComponentImageRepository	controllers/component_image_controller.go:299	Component updated successfully	{"controller": "component", "controllerGroup": "appstudio.redhat.com", "controllerKind": "Component", "Component": {"name":"build-suite-test-component-image-source-ajdr","namespace":"build-e2e-rsql-tenant"}, "namespace": "build-e2e-rsql-tenant", "name": "build-suite-test-component-image-source-ajdr", "reconcileID": "1c1f7548-b16a-43ba-a91b-0a2aa32cc6cd", "action": "UPDATE"}
2023-08-18T01:19:57.257Z	INFO	ComponentImageRepository	controllers/component_image_controller.go:170	Waiting for devfile model in component	{"controller": "component", "controllerGroup": "appstudio.redhat.com", "controllerKind": "Component", "Component": {"name":"devfile-sample-hello-world-0dwf","namespace":"build-e2e-bslz-tenant"}, "namespace": "build-e2e-bslz-tenant", "name": "devfile-sample-hello-world-0dwf", "reconcileID": "188862ec-a820-473e-a3c1-a2bb87031138"}
2023-08-18T01:19:57.327Z	INFO	ComponentImageRepository	controllers/component_image_controller.go:170	Waiting for devfile model in component	{"controller": "component", "controllerGroup": "appstudio.redhat.com", "controllerKind": "Component", "Component": {"name":"devfile-sample-hello-world-0dwf","namespace":"build-e2e-bslz-tenant"}, "namespace": "build-e2e-bslz-tenant", "name": "devfile-sample-hello-world-0dwf", "reconcileID": "88549a15-0c83-4ac0-9e2d-f754e4aea532"}`, filtered)
}

func TestJsonLogParsing(t *testing.T) {

	start, _ := time.Parse(time.RFC3339, "2023-08-18T01:30:06Z")
	filtered := FilterLogs(jsonLogs, start)
	assert.Equal(t,
		`{"level":"info","ts":"2023-08-18T01:30:06Z","logger":"artifactbuild","caller":"artifactbuild/artifactbuild.go:234","msg":"ArtifactBuild shaded.jdk11.1.9-c65abf6b changing state from ArtifactBuildDiscovering to ArtifactBuildBuilding","namespace":"rhtap-demo-afcg-tenant","resource":"shaded.jdk11.1.9-c65abf6b","kind":"ArtifactBuild","ab-gav":"io.github.stuartwdouglas.hacbs-test.shaded:shaded-jdk11:1.9","ab-initial-state":"ArtifactBuildDiscovering"}
{"level":"info","ts":"2023-08-18T01:31:06Z","logger":"artifactbuild","caller":"artifactbuild/artifactbuild.go:234","msg":"ArtifactBuild shaded.jdk11.1.9-c65abf6b changing state from ArtifactBuildBuilding to ArtifactBuildNew","namespace":"rhtap-demo-afcg-tenant","resource":"shaded.jdk11.1.9-c65abf6b","kind":"ArtifactBuild","ab-gav":"io.github.stuartwdouglas.hacbs-test.shaded:shaded-jdk11:1.9","ab-initial-state":"ArtifactBuildBuilding"}
{"level":"info","ts":"2023-08-18T01:32:06Z","logger":"artifactbuild","caller":"artifactbuild/artifactbuild.go:524","msg":"Found pipeline run with community dependencies","namespace":"rhtap-demo-afcg-tenant","resource":"hacbs-test-project-jyxg-on-push-vxwtr","kind":"PipelineRun"}
{"level":"info","ts":"2023-08-18T01:33:06Z","logger":"artifactbuild","caller":"artifactbuild/artifactbuild.go:530","msg":"Found community dependency, creating ArtifactBuild","namespace":"rhtap-demo-afcg-tenant","resource":"hacbs-test-project-jyxg-on-push-vxwtr","kind":"PipelineRun","gav":"io.github.stuartwdouglas.hacbs-test.gradle:hacbs-test-simple-gradle-jdk8:1.1","artifactbuild":"hacbs.test.simple.gradle.jdk8.1.1-f6a4dce0","action":"ADD"}
{"level":"info","ts":"2023-08-18T01:34:06Z","logger":"artifactbuild","caller":"artifactbuild/artifactbuild.go:530","msg":"Found community dependency, creating ArtifactBuild","namespace":"rhtap-demo-afcg-tenant","resource":"hacbs-test-project-jyxg-on-push-vxwtr","kind":"PipelineRun","gav":"io.github.stuartwdouglas.hacbs-test.shaded:shaded-jdk11:1.9","artifactbuild":"shaded.jdk11.1.9-c65abf6b","action":"ADD"}
{"level":"info","ts":"2023-08-18T01:35:06Z","logger":"artifactbuild","caller":"artifactbuild/artifactbuild.go:530","msg":"Found community dependency, creating ArtifactBuild","namespace":"rhtap-demo-afcg-tenant","resource":"hacbs-test-project-jyxg-on-push-vxwtr","kind":"PipelineRun","gav":"io.github.stuartwdouglas.hacbs-test.simple:simple-jdk17:0.1.2","artifactbuild":"simple.jdk17.0.1.2-22fafbfd","action":"ADD"}`, filtered)
}