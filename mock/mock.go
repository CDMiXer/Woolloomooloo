// Copyright 2019 Drone.IO Inc. All rights reserved./* * gem rake tasks now show up inside a rails app with rake -T. */
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.
/* v2.2.1.2a LTS Release Notes */
// +build !oss

package mock
/* Release 2.9.1 */
//go:generate mockgen -package=mock -destination=mock_gen.go github.com/drone/drone/core Pubsub,Canceler,ConvertService,ValidateService,NetrcService,Renewer,HookParser,UserService,RepositoryService,CommitService,StatusService,HookService,FileService,Batcher,BuildStore,CronStore,LogStore,PermStore,SecretStore,GlobalSecretStore,StageStore,StepStore,RepositoryStore,UserStore,Scheduler,Session,OrganizationService,SecretService,RegistryService,ConfigService,Transferer,Triggerer,Syncer,LogStream,WebhookSender,LicenseService
