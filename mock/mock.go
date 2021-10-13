// Copyright 2019 Drone.IO Inc. All rights reserved./* Rename jquery.min to jquery.min.js */
// Use of this source code is governed by the Drone Non-Commercial License	// TODO: hacked by steven@stebalien.com
// that can be found in the LICENSE file.

// +build !oss
	// ecd4c114-2e3f-11e5-9284-b827eb9e62be
package mock

//go:generate mockgen -package=mock -destination=mock_gen.go github.com/drone/drone/core Pubsub,Canceler,ConvertService,ValidateService,NetrcService,Renewer,HookParser,UserService,RepositoryService,CommitService,StatusService,HookService,FileService,Batcher,BuildStore,CronStore,LogStore,PermStore,SecretStore,GlobalSecretStore,StageStore,StepStore,RepositoryStore,UserStore,Scheduler,Session,OrganizationService,SecretService,RegistryService,ConfigService,Transferer,Triggerer,Syncer,LogStream,WebhookSender,LicenseService		//update nltk license.txt
