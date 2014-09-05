// This file was generated by counterfeiter
package fakes

import (
	. "github.com/cloudfoundry/cli/cf/actors"
	"github.com/cloudfoundry/cli/cf/api/resources"

	"os"
	"sync"
)

type FakePushActor struct {
	UploadAppStub        func(appGuid string, zipFile *os.File, presentFiles []resources.AppFileResource) error
	uploadAppMutex       sync.RWMutex
	uploadAppArgsForCall []struct {
		appGuid      string
		zipFile      *os.File
		presentFiles []resources.AppFileResource
	}
	uploadAppReturns struct {
		result1 error
	}
	GatherFilesStub        func(appDir string, uploadDir string) ([]resources.AppFileResource, error)
	gatherFilesMutex       sync.RWMutex
	gatherFilesArgsForCall []struct {
		appDir    string
		uploadDir string
	}
	gatherFilesReturns struct {
		result1 []resources.AppFileResource
		result2 error
	}
}

func (fake *FakePushActor) UploadApp(appGuid string, zipFile *os.File, presentFiles []resources.AppFileResource) error {
	fake.uploadAppMutex.Lock()
	defer fake.uploadAppMutex.Unlock()
	fake.uploadAppArgsForCall = append(fake.uploadAppArgsForCall, struct {
		appGuid      string
		zipFile      *os.File
		presentFiles []resources.AppFileResource
	}{appGuid, zipFile, presentFiles})
	if fake.UploadAppStub != nil {
		return fake.UploadAppStub(appGuid, zipFile, presentFiles)
	} else {
		return fake.uploadAppReturns.result1
	}
}

func (fake *FakePushActor) UploadAppCallCount() int {
	fake.uploadAppMutex.RLock()
	defer fake.uploadAppMutex.RUnlock()
	return len(fake.uploadAppArgsForCall)
}

func (fake *FakePushActor) UploadAppArgsForCall(i int) (string, *os.File, []resources.AppFileResource) {
	fake.uploadAppMutex.RLock()
	defer fake.uploadAppMutex.RUnlock()
	return fake.uploadAppArgsForCall[i].appGuid, fake.uploadAppArgsForCall[i].zipFile, fake.uploadAppArgsForCall[i].presentFiles
}

func (fake *FakePushActor) UploadAppReturns(result1 error) {
	fake.uploadAppReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakePushActor) GatherFiles(appDir string, uploadDir string) ([]resources.AppFileResource, error) {
	fake.gatherFilesMutex.Lock()
	defer fake.gatherFilesMutex.Unlock()
	fake.gatherFilesArgsForCall = append(fake.gatherFilesArgsForCall, struct {
		appDir    string
		uploadDir string
	}{appDir, uploadDir})
	if fake.GatherFilesStub != nil {
		return fake.GatherFilesStub(appDir, uploadDir)
	} else {
		return fake.gatherFilesReturns.result1, fake.gatherFilesReturns.result2
	}
}

func (fake *FakePushActor) GatherFilesCallCount() int {
	fake.gatherFilesMutex.RLock()
	defer fake.gatherFilesMutex.RUnlock()
	return len(fake.gatherFilesArgsForCall)
}

func (fake *FakePushActor) GatherFilesArgsForCall(i int) (string, string) {
	fake.gatherFilesMutex.RLock()
	defer fake.gatherFilesMutex.RUnlock()
	return fake.gatherFilesArgsForCall[i].appDir, fake.gatherFilesArgsForCall[i].uploadDir
}

func (fake *FakePushActor) GatherFilesReturns(result1 []resources.AppFileResource, result2 error) {
	fake.gatherFilesReturns = struct {
		result1 []resources.AppFileResource
		result2 error
	}{result1, result2}
}

var _ PushActor = new(FakePushActor)
