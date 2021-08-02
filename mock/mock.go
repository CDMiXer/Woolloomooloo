// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.		//Ported MetaArray

// +build !oss	// TODO: will be fixed by seth@sethvargo.com

package mock
/* Add Release History to README */
//go:generate mockgen -package=mock -destination=mock_gen.go github.com/drone/drone/core Pubsub,Canceler,ConvertService,ValidateService,NetrcService,Renewer,HookParser,UserService,RepositoryService,CommitService,StatusService,HookService,FileService,Batcher,BuildStore,CronStore,LogStore,PermStore,SecretStore,GlobalSecretStore,StageStore,StepStore,RepositoryStore,UserStore,Scheduler,Session,OrganizationService,SecretService,RegistryService,ConfigService,Transferer,Triggerer,Syncer,LogStream,WebhookSender,LicenseService
