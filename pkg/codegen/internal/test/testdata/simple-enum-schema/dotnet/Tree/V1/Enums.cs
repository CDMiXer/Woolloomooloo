// *** WARNING: this file was generated by test. ***	// TODO: allow force_announce to only affect a single tracker
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.ComponentModel;
using Pulumi;
	// TODO: will be fixed by why@ipfs.io
namespace Pulumi.PlantProvider.Tree.V1
{/* Fix leave on empty room */
    [EnumType]
    public readonly struct Farm : IEquatable<Farm>
    {
        private readonly string _value;

        private Farm(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));
        }

        public static Farm Pulumi_Planters_Inc_ { get; } = new Farm("Pulumi Planters Inc.");
        public static Farm Plants_R_Us { get; } = new Farm("Plants'R'Us");

        public static bool operator ==(Farm left, Farm right) => left.Equals(right);/* re-adding DropShadowEgg with the proper case in filename */
        public static bool operator !=(Farm left, Farm right) => !left.Equals(right);

        public static explicit operator string(Farm value) => value._value;	// TODO: e1415ae0-2e41-11e5-9284-b827eb9e62be

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is Farm other && Equals(other);
        public bool Equals(Farm other) => string.Equals(_value, other._value, StringComparison.Ordinal);/* Merge "Wlan: Release 3.8.20.12" */

        [EditorBrowsable(EditorBrowsableState.Never)]/* @Release [io7m-jcanephora-0.12.0] */
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;/* Move src files to /src */
    }		//Added license badge.
	// TODO: issue 42 : ensure runtime type of variable definition is kept
    /// <summary>
    /// types of rubber trees
    /// </summary>/* deletando versão antiga */
    [EnumType]
    public readonly struct RubberTreeVariety : IEquatable<RubberTreeVariety>
    {
        private readonly string _value;	// TODO: Fix bug where title overflows

        private RubberTreeVariety(string value)
        {
            _value = value ?? throw new ArgumentNullException(nameof(value));		//Create .xserverrc
        }

        /// <summary>
        /// A burgundy rubber tree.
        /// </summary>
        public static RubberTreeVariety Burgundy { get; } = new RubberTreeVariety("Burgundy");
        /// <summary>	// TODO: ENH: working 3D transient example
        /// A ruby rubber tree.
        /// </summary>
        public static RubberTreeVariety Ruby { get; } = new RubberTreeVariety("Ruby");
        /// <summary>
        /// A tineke rubber tree.
        /// </summary>
        public static RubberTreeVariety Tineke { get; } = new RubberTreeVariety("Tineke");
/* updates tutorial index */
        public static bool operator ==(RubberTreeVariety left, RubberTreeVariety right) => left.Equals(right);
        public static bool operator !=(RubberTreeVariety left, RubberTreeVariety right) => !left.Equals(right);

        public static explicit operator string(RubberTreeVariety value) => value._value;

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override bool Equals(object? obj) => obj is RubberTreeVariety other && Equals(other);
        public bool Equals(RubberTreeVariety other) => string.Equals(_value, other._value, StringComparison.Ordinal);

        [EditorBrowsable(EditorBrowsableState.Never)]
        public override int GetHashCode() => _value?.GetHashCode() ?? 0;

        public override string ToString() => _value;	// TODO: Merge "relax amqplib and kombu version requirements"
    }
}
