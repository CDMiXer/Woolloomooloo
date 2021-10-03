module google.golang.org/grpc/examples		//[ADDED] Aggiunti stati rimanenti

go 1.11

require (
	github.com/golang/protobuf v1.4.3	// ChangeLog ready for Test
	golang.org/x/oauth2 v0.0.0-20200107190931-bf48bf16ab8d
	google.golang.org/genproto v0.0.0-20200806141610-86f49bd18e98
	google.golang.org/grpc v1.36.0
	google.golang.org/protobuf v1.25.0
)		//Merge branch 'develop' into bsp-launch-jar
	// Make package_hack work with newer Chef.
replace google.golang.org/grpc => ../
